package querier

import (
	"flag"
	"fmt"
	"sync"

	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmselect/netstorage"
	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmselect/searchutils"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/auth"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logql"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/storage"
)

var (
	disableCache           = flag.Bool("search.disableCache", false, "Whether to disable response caching. This may be useful during data backfilling")
	maxPointsPerTimeseries = flag.Int("search.maxPointsPerTimeseries", 30e3, "The maximum points per a single timeseries returned from the search")
)

// The minimum number of points per timeseries for enabling time rounding.
// This improves cache hit ratio for frequently requested queries over
// big time ranges.
const minTimeseriesPointsForTimeRounding = 50

// ValidateMaxPointsPerTimeseries checks the maximum number of points that
// may be returned per each time series.
//
// The number mustn't exceed -search.maxPointsPerTimeseries.
func ValidateMaxPointsPerTimeseries(start, end, step int64) error {
	points := (end-start)/step + 1
	if uint64(points) > uint64(*maxPointsPerTimeseries) {
		return fmt.Errorf(`too many points for the given step=%d, start=%d and end=%d: %d; cannot exceed -search.maxPointsPerTimeseries=%d`,
			step, start, end, uint64(points), *maxPointsPerTimeseries)
	}
	return nil
}

// AdjustStartEnd adjusts start and end values, so response caching may be enabled.
//
// See EvalConfig.mayCache for details.
func AdjustStartEnd(start, end, step int64) (int64, int64) {
	if *disableCache {
		// Do not adjust start and end values when cache is disabled.
		// See https://github.com/VictoriaMetrics/VictoriaMetrics/issues/563
		return start, end
	}
	points := (end-start)/step + 1
	if points < minTimeseriesPointsForTimeRounding {
		// Too small number of points for rounding.
		return start, end
	}

	// Round start and end to values divisible by step in order
	// to enable response caching (see EvalConfig.mayCache).
	start, end = alignStartEnd(start, end, step)

	// Make sure that the new number of points is the same as the initial number of points.
	newPoints := (end-start)/step + 1
	for newPoints > points {
		end -= step
		newPoints--
	}

	return start, end
}

func alignStartEnd(start, end, step int64) (int64, int64) {
	// Round start to the nearest smaller value divisible by step.
	start -= start % step
	// Round end to the nearest bigger value divisible by step.
	adjust := end % step
	if adjust > 0 {
		end += step - adjust
	}
	return start, end
}

// EvalConfig is the configuration required for query evaluation via Exec
type EvalConfig struct {
	AuthToken *auth.Token
	Start     int64
	End       int64
	Step      int64

	// QuotedRemoteAddr contains quoted remote address.
	QuotedRemoteAddr string

	Deadline searchutils.Deadline

	MayCache bool

	// LookbackDelta is analog to `-query.lookback-delta` from Prometheus.
	LookbackDelta int64

	DenyPartialResponse bool

	timestamps     []int64
	timestampsOnce sync.Once
}

func (ec *EvalConfig) validate() {
	if ec.Start > ec.End {
		logger.Panicf("BUG: start cannot exceed end; got %d vs %d", ec.Start, ec.End)
	}
	if ec.Step <= 0 {
		logger.Panicf("BUG: step must be greater than 0; got %d", ec.Step)
	}
}

func (ec *EvalConfig) mayCache() bool {
	if *disableCache {
		return false
	}
	if !ec.MayCache {
		return false
	}
	if ec.Start%ec.Step != 0 {
		return false
	}
	if ec.End%ec.Step != 0 {
		return false
	}
	return true
}

func (ec *EvalConfig) getSharedTimestamps() []int64 {
	ec.timestampsOnce.Do(ec.timestampsInit)
	return ec.timestamps
}

func (ec *EvalConfig) timestampsInit() {
	ec.timestamps = getTimestamps(ec.Start, ec.End, ec.Step)
}

func getTimestamps(start, end, step int64) []int64 {
	// Sanity checks.
	if step <= 0 {
		logger.Panicf("BUG: Step must be bigger than 0; got %d", step)
	}
	if start > end {
		logger.Panicf("BUG: Start cannot exceed End; got %d vs %d", start, end)
	}
	if err := ValidateMaxPointsPerTimeseries(start, end, step); err != nil {
		logger.Panicf("BUG: %s; this must be validated before the call to getTimestamps", err)
	}

	// Prepare timestamps.
	points := 1 + (end-start)/step
	timestamps := make([]int64, points)
	for i := range timestamps {
		timestamps[i] = start
		start += step
	}
	return timestamps
}

func evalExpr(ec *EvalConfig, e logql.Expr) ([]*timeseries, error) {
	if me, ok := e.(*logql.MetricExpr); ok {
		rv, err := evalMetricExpr(ec, me)
		if err != nil {
			return nil, fmt.Errorf(`cannot evaluate %q: %w`, me.AppendString(nil), err)
		}
		return rv, nil
	}
	return nil, fmt.Errorf("unexpected expression %q", e.AppendString(nil))
}

var nan []byte = nil

func evalMetricExpr(ec *EvalConfig, me *logql.MetricExpr) ([]*timeseries, error) {
	if me.IsEmpty() {
		return evalNumber(ec, nan), nil
	}

	// Obtain rollup configs before fetching data from db,
	// so type errors can be caught earlier.
	sharedTimestamps := getTimestamps(ec.Start, ec.End, ec.Step)

	// Fetch the remaining part of the result.
	tfs := toTagFilters(me.LabelFilters)
	minTimestamp := ec.Start - maxSilenceInterval - ec.Step

	sq := &storage.SearchQuery{
		AccountID:    ec.AuthToken.AccountID,
		ProjectID:    ec.AuthToken.ProjectID,
		MinTimestamp: minTimestamp,
		MaxTimestamp: ec.End,
		TagFilterss:  [][]storage.TagFilter{tfs},
	}
	rss, isPartial, err := netstorage.ProcessSearchQuery(ec.AuthToken, sq, true, ec.Deadline)
	if err != nil {
		return nil, err
	}
	if isPartial && ec.DenyPartialResponse {
		rss.Cancel()
		return nil, fmt.Errorf("cannot return full response, since some of vmstorage nodes are unavailable")
	}
	rssLen := rss.Len()
	if rssLen == 0 {
		rss.Cancel()
		return nil, nil
	}

	// Evaluate rollup
	var tss []*timeseries
	var tssLock sync.Mutex
	err = rss.RunParallel(func(rs *netstorage.Result, workerID uint) error {
		var ts timeseries

		ts.MetricName.CopyFrom(&rs.MetricName)
		ts.Values = rs.Values
		ts.Timestamps = sharedTimestamps
		ts.denyReuse = true

		tssLock.Lock()
		tss = append(tss, &ts)
		tssLock.Unlock()
		return nil
	})
	if err != nil {
		return nil, err
	}

	return tss, nil
}

func evalNumber(ec *EvalConfig, n []byte) []*timeseries {
	var ts timeseries
	ts.denyReuse = true
	ts.MetricName.AccountID = ec.AuthToken.AccountID
	ts.MetricName.ProjectID = ec.AuthToken.ProjectID
	timestamps := ec.getSharedTimestamps()
	values := make([][]byte, len(timestamps))
	for i := range timestamps {
		values[i] = n
	}
	ts.Values = values
	ts.Timestamps = timestamps
	return []*timeseries{&ts}
}

func toTagFilters(lfs []logql.LabelFilter) []storage.TagFilter {
	tfs := make([]storage.TagFilter, len(lfs))
	for i := range lfs {
		toTagFilter(&tfs[i], &lfs[i])
	}
	return tfs
}

func toTagFilter(dst *storage.TagFilter, src *logql.LabelFilter) {
	if src.Label != "__name__" {
		dst.Key = []byte(src.Label)
	} else {
		// This is required for storage.Search.
		dst.Key = nil
	}
	dst.Value = []byte(src.Value)
	dst.IsRegexp = src.IsRegexp
	dst.IsNegative = src.IsNegative
}