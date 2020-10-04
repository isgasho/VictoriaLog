// Code generated by qtc from "export.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmselect/prometheus/export.qtpl:1
package prometheus

//line app/vmselect/prometheus/export.qtpl:1
import (
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/storage"
	"github.com/valyala/quicktemplate"
)

//line app/vmselect/prometheus/export.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/prometheus/export.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/prometheus/export.qtpl:8
func StreamExportPrometheusLine(qw422016 *qt422016.Writer, xb *exportBlock) {
//line app/vmselect/prometheus/export.qtpl:9
	if len(xb.timestamps) == 0 {
//line app/vmselect/prometheus/export.qtpl:9
		return
//line app/vmselect/prometheus/export.qtpl:9
	}
//line app/vmselect/prometheus/export.qtpl:10
	bb := quicktemplate.AcquireByteBuffer()

//line app/vmselect/prometheus/export.qtpl:11
	writeprometheusMetricName(bb, xb.mn)

//line app/vmselect/prometheus/export.qtpl:12
	for i, ts := range xb.timestamps {
//line app/vmselect/prometheus/export.qtpl:13
		qw422016.N().Z(bb.B)
//line app/vmselect/prometheus/export.qtpl:13
		qw422016.N().S(` `)
//line app/vmselect/prometheus/export.qtpl:14
		qw422016.N().Z(xb.datas[i])
//line app/vmselect/prometheus/export.qtpl:14
		qw422016.N().S(` `)
//line app/vmselect/prometheus/export.qtpl:15
		qw422016.N().DL(ts)
//line app/vmselect/prometheus/export.qtpl:15
		qw422016.N().S(`
`)
//line app/vmselect/prometheus/export.qtpl:16
	}
//line app/vmselect/prometheus/export.qtpl:17
	quicktemplate.ReleaseByteBuffer(bb)

//line app/vmselect/prometheus/export.qtpl:18
}

//line app/vmselect/prometheus/export.qtpl:18
func WriteExportPrometheusLine(qq422016 qtio422016.Writer, xb *exportBlock) {
//line app/vmselect/prometheus/export.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/export.qtpl:18
	StreamExportPrometheusLine(qw422016, xb)
//line app/vmselect/prometheus/export.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/export.qtpl:18
}

//line app/vmselect/prometheus/export.qtpl:18
func ExportPrometheusLine(xb *exportBlock) string {
//line app/vmselect/prometheus/export.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/export.qtpl:18
	WriteExportPrometheusLine(qb422016, xb)
//line app/vmselect/prometheus/export.qtpl:18
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/export.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/export.qtpl:18
	return qs422016
//line app/vmselect/prometheus/export.qtpl:18
}

//line app/vmselect/prometheus/export.qtpl:20
func StreamExportJSONLine(qw422016 *qt422016.Writer, xb *exportBlock) {
//line app/vmselect/prometheus/export.qtpl:21
	if len(xb.timestamps) == 0 {
//line app/vmselect/prometheus/export.qtpl:21
		return
//line app/vmselect/prometheus/export.qtpl:21
	}
//line app/vmselect/prometheus/export.qtpl:21
	qw422016.N().S(`{"metric":`)
//line app/vmselect/prometheus/export.qtpl:23
	streammetricNameObject(qw422016, xb.mn)
//line app/vmselect/prometheus/export.qtpl:23
	qw422016.N().S(`,"values":[`)
//line app/vmselect/prometheus/export.qtpl:25
	if len(xb.datas) > 0 {
//line app/vmselect/prometheus/export.qtpl:26
		values := xb.datas

//line app/vmselect/prometheus/export.qtpl:27
		qw422016.N().Z(values[0])
//line app/vmselect/prometheus/export.qtpl:28
		values = values[1:]

//line app/vmselect/prometheus/export.qtpl:29
		for _, v := range values {
//line app/vmselect/prometheus/export.qtpl:29
			qw422016.N().S(`,`)
//line app/vmselect/prometheus/export.qtpl:30
			qw422016.N().Z(v)
//line app/vmselect/prometheus/export.qtpl:31
		}
//line app/vmselect/prometheus/export.qtpl:32
	}
//line app/vmselect/prometheus/export.qtpl:32
	qw422016.N().S(`],"timestamps":[`)
//line app/vmselect/prometheus/export.qtpl:35
	if len(xb.timestamps) > 0 {
//line app/vmselect/prometheus/export.qtpl:36
		timestamps := xb.timestamps

//line app/vmselect/prometheus/export.qtpl:37
		qw422016.N().DL(timestamps[0])
//line app/vmselect/prometheus/export.qtpl:38
		timestamps = timestamps[1:]

//line app/vmselect/prometheus/export.qtpl:39
		for _, ts := range timestamps {
//line app/vmselect/prometheus/export.qtpl:39
			qw422016.N().S(`,`)
//line app/vmselect/prometheus/export.qtpl:40
			qw422016.N().DL(ts)
//line app/vmselect/prometheus/export.qtpl:41
		}
//line app/vmselect/prometheus/export.qtpl:42
	}
//line app/vmselect/prometheus/export.qtpl:42
	qw422016.N().S(`]}`)
//line app/vmselect/prometheus/export.qtpl:44
	qw422016.N().S(`
`)
//line app/vmselect/prometheus/export.qtpl:45
}

//line app/vmselect/prometheus/export.qtpl:45
func WriteExportJSONLine(qq422016 qtio422016.Writer, xb *exportBlock) {
//line app/vmselect/prometheus/export.qtpl:45
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/export.qtpl:45
	StreamExportJSONLine(qw422016, xb)
//line app/vmselect/prometheus/export.qtpl:45
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/export.qtpl:45
}

//line app/vmselect/prometheus/export.qtpl:45
func ExportJSONLine(xb *exportBlock) string {
//line app/vmselect/prometheus/export.qtpl:45
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/export.qtpl:45
	WriteExportJSONLine(qb422016, xb)
//line app/vmselect/prometheus/export.qtpl:45
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/export.qtpl:45
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/export.qtpl:45
	return qs422016
//line app/vmselect/prometheus/export.qtpl:45
}

//line app/vmselect/prometheus/export.qtpl:47
func StreamExportPromAPILine(qw422016 *qt422016.Writer, xb *exportBlock) {
//line app/vmselect/prometheus/export.qtpl:47
	qw422016.N().S(`{"metric":`)
//line app/vmselect/prometheus/export.qtpl:49
	streammetricNameObject(qw422016, xb.mn)
//line app/vmselect/prometheus/export.qtpl:49
	qw422016.N().S(`,"values":`)
//line app/vmselect/prometheus/export.qtpl:50
	streamdatasWithTimestamps(qw422016, xb.datas, xb.timestamps)
//line app/vmselect/prometheus/export.qtpl:50
	qw422016.N().S(`}`)
//line app/vmselect/prometheus/export.qtpl:52
}

//line app/vmselect/prometheus/export.qtpl:52
func WriteExportPromAPILine(qq422016 qtio422016.Writer, xb *exportBlock) {
//line app/vmselect/prometheus/export.qtpl:52
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/export.qtpl:52
	StreamExportPromAPILine(qw422016, xb)
//line app/vmselect/prometheus/export.qtpl:52
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/export.qtpl:52
}

//line app/vmselect/prometheus/export.qtpl:52
func ExportPromAPILine(xb *exportBlock) string {
//line app/vmselect/prometheus/export.qtpl:52
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/export.qtpl:52
	WriteExportPromAPILine(qb422016, xb)
//line app/vmselect/prometheus/export.qtpl:52
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/export.qtpl:52
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/export.qtpl:52
	return qs422016
//line app/vmselect/prometheus/export.qtpl:52
}

//line app/vmselect/prometheus/export.qtpl:54
func StreamExportPromAPIResponse(qw422016 *qt422016.Writer, resultsCh <-chan *quicktemplate.ByteBuffer) {
//line app/vmselect/prometheus/export.qtpl:54
	qw422016.N().S(`{"status":"success","data":{"resultType":"matrix","result":[`)
//line app/vmselect/prometheus/export.qtpl:60
	bb, ok := <-resultsCh

//line app/vmselect/prometheus/export.qtpl:61
	if ok {
//line app/vmselect/prometheus/export.qtpl:62
		qw422016.N().Z(bb.B)
//line app/vmselect/prometheus/export.qtpl:63
		quicktemplate.ReleaseByteBuffer(bb)

//line app/vmselect/prometheus/export.qtpl:64
		for bb := range resultsCh {
//line app/vmselect/prometheus/export.qtpl:64
			qw422016.N().S(`,`)
//line app/vmselect/prometheus/export.qtpl:65
			qw422016.N().Z(bb.B)
//line app/vmselect/prometheus/export.qtpl:66
			quicktemplate.ReleaseByteBuffer(bb)

//line app/vmselect/prometheus/export.qtpl:67
		}
//line app/vmselect/prometheus/export.qtpl:68
	}
//line app/vmselect/prometheus/export.qtpl:68
	qw422016.N().S(`]}}`)
//line app/vmselect/prometheus/export.qtpl:72
}

//line app/vmselect/prometheus/export.qtpl:72
func WriteExportPromAPIResponse(qq422016 qtio422016.Writer, resultsCh <-chan *quicktemplate.ByteBuffer) {
//line app/vmselect/prometheus/export.qtpl:72
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/export.qtpl:72
	StreamExportPromAPIResponse(qw422016, resultsCh)
//line app/vmselect/prometheus/export.qtpl:72
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/export.qtpl:72
}

//line app/vmselect/prometheus/export.qtpl:72
func ExportPromAPIResponse(resultsCh <-chan *quicktemplate.ByteBuffer) string {
//line app/vmselect/prometheus/export.qtpl:72
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/export.qtpl:72
	WriteExportPromAPIResponse(qb422016, resultsCh)
//line app/vmselect/prometheus/export.qtpl:72
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/export.qtpl:72
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/export.qtpl:72
	return qs422016
//line app/vmselect/prometheus/export.qtpl:72
}

//line app/vmselect/prometheus/export.qtpl:74
func StreamExportStdResponse(qw422016 *qt422016.Writer, resultsCh <-chan *quicktemplate.ByteBuffer) {
//line app/vmselect/prometheus/export.qtpl:75
	for bb := range resultsCh {
//line app/vmselect/prometheus/export.qtpl:76
		qw422016.N().Z(bb.B)
//line app/vmselect/prometheus/export.qtpl:77
		quicktemplate.ReleaseByteBuffer(bb)

//line app/vmselect/prometheus/export.qtpl:78
	}
//line app/vmselect/prometheus/export.qtpl:79
}

//line app/vmselect/prometheus/export.qtpl:79
func WriteExportStdResponse(qq422016 qtio422016.Writer, resultsCh <-chan *quicktemplate.ByteBuffer) {
//line app/vmselect/prometheus/export.qtpl:79
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/export.qtpl:79
	StreamExportStdResponse(qw422016, resultsCh)
//line app/vmselect/prometheus/export.qtpl:79
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/export.qtpl:79
}

//line app/vmselect/prometheus/export.qtpl:79
func ExportStdResponse(resultsCh <-chan *quicktemplate.ByteBuffer) string {
//line app/vmselect/prometheus/export.qtpl:79
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/export.qtpl:79
	WriteExportStdResponse(qb422016, resultsCh)
//line app/vmselect/prometheus/export.qtpl:79
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/export.qtpl:79
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/export.qtpl:79
	return qs422016
//line app/vmselect/prometheus/export.qtpl:79
}

//line app/vmselect/prometheus/export.qtpl:81
func streamprometheusMetricName(qw422016 *qt422016.Writer, mn *storage.MetricName) {
//line app/vmselect/prometheus/export.qtpl:82
	qw422016.N().Z(mn.MetricGroup)
//line app/vmselect/prometheus/export.qtpl:83
	if len(mn.Tags) > 0 {
//line app/vmselect/prometheus/export.qtpl:83
		qw422016.N().S(`{`)
//line app/vmselect/prometheus/export.qtpl:85
		tags := mn.Tags

//line app/vmselect/prometheus/export.qtpl:86
		qw422016.N().Z(tags[0].Key)
//line app/vmselect/prometheus/export.qtpl:86
		qw422016.N().S(`=`)
//line app/vmselect/prometheus/export.qtpl:86
		qw422016.N().QZ(tags[0].Value)
//line app/vmselect/prometheus/export.qtpl:87
		tags = tags[1:]

//line app/vmselect/prometheus/export.qtpl:88
		for i := range tags {
//line app/vmselect/prometheus/export.qtpl:89
			tag := &tags[i]

//line app/vmselect/prometheus/export.qtpl:89
			qw422016.N().S(`,`)
//line app/vmselect/prometheus/export.qtpl:90
			qw422016.N().Z(tag.Key)
//line app/vmselect/prometheus/export.qtpl:90
			qw422016.N().S(`=`)
//line app/vmselect/prometheus/export.qtpl:90
			qw422016.N().QZ(tag.Value)
//line app/vmselect/prometheus/export.qtpl:91
		}
//line app/vmselect/prometheus/export.qtpl:91
		qw422016.N().S(`}`)
//line app/vmselect/prometheus/export.qtpl:93
	}
//line app/vmselect/prometheus/export.qtpl:94
}

//line app/vmselect/prometheus/export.qtpl:94
func writeprometheusMetricName(qq422016 qtio422016.Writer, mn *storage.MetricName) {
//line app/vmselect/prometheus/export.qtpl:94
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/export.qtpl:94
	streamprometheusMetricName(qw422016, mn)
//line app/vmselect/prometheus/export.qtpl:94
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/export.qtpl:94
}

//line app/vmselect/prometheus/export.qtpl:94
func prometheusMetricName(mn *storage.MetricName) string {
//line app/vmselect/prometheus/export.qtpl:94
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/export.qtpl:94
	writeprometheusMetricName(qb422016, mn)
//line app/vmselect/prometheus/export.qtpl:94
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/export.qtpl:94
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/export.qtpl:94
	return qs422016
//line app/vmselect/prometheus/export.qtpl:94
}
