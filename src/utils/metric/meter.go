package metric

import (
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"testTools/src/utils/clog"

	"math"
	"sync"
	"time"
)

var (
	ErrorMetricNotFound = errors.New("metric.not.found")
)

const (
	CurCounterSuffix  = "_cur"
	LastCounterSuffix = "_last"

	NilMetricValue = math.SmallestNonzeroFloat64

	DefaultGCPeriod       = 30 // unit: minutes
	DefaultMaxMetricCount = 10000
)

type Meter interface {
	// Inc increments the Meter by 1. Use Add to increment it by arbitrary
	// values.
	Inc(prometheus.Labels)
	// Add adds the given value to the meter. It panics if the value is <
	// 0.
	Add(prometheus.Labels, float64)
	//
	GetName() string
}

type MeterOpts prometheus.Opts

type meter struct {
	interval int // statistical period
	lock     sync.Mutex

	name           string
	help           string
	variableLabels []string

	gauge     *prometheus.GaugeVec
	last, cur *prometheus.CounterVec

	lastGCTime time.Time
}

//func NewMeter(opt MeterOpts) Meter {
//
//}

//
// example:
// scriptExecQPSMetric := NewMeterVec(10, MeterOpts{
//		Name: ScriptExecuteQPS,
//		Help: "qps of script execute",
//	},
//		[]string{LabelService, LabelRelease, LabelSpace},)
// scriptExecQPSMetric.Inc(prometheus.Labels{LabelService: "jvessel", LabelRelease: "0.0.1", LabelSpace: "i-52ei4vxsp0"})
// scriptExecQPSMetric.Add(prometheus.Labels{LabelService: "jvessel", LabelRelease: "0.0.1", LabelSpace: "i-52ei4vxsp0"}, 10.5)
//

func NewMeterVec(interval int, opts MeterOpts, labelNames []string) Meter {
	curOpts, lastOpts := opts, opts
	curOpts.Name = opts.Name + CurCounterSuffix
	lastOpts.Name = opts.Name + LastCounterSuffix
	m := &meter{
		interval:       interval,
		name:           opts.Name,
		help:           opts.Name,
		variableLabels: labelNames,
		lastGCTime:     time.Now(),
		gauge:          prometheus.NewGaugeVec(prometheus.GaugeOpts(opts), labelNames),
		cur:            prometheus.NewCounterVec(prometheus.CounterOpts(curOpts), labelNames),
		last:           prometheus.NewCounterVec(prometheus.CounterOpts(lastOpts), labelNames),
	}
	// register metric
	prometheus.MustRegister(m.cur)
	prometheus.MustRegister(m.last)
	prometheus.MustRegister(m.gauge)

	go m.tick()

	return m
}

func (m *meter) Inc(labels prometheus.Labels) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.cur.With(labels).Inc()
}

func (m *meter) Add(labels prometheus.Labels, v float64) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.cur.With(labels).Add(v)
}

func (m *meter) GetName() string {
	return m.name
}

func (m *meter) getMetricFamily(metricType dto.MetricType, metricName string) *dto.MetricFamily {
	metricFamilies, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		clog.Blog.Warningf("get metric family from default gatherer failed. Error: %v", err)
		return nil
	}

	for _, metricFamily := range metricFamilies {
		if *metricFamily.Type == metricType && *metricFamily.Name == metricName {
			return metricFamily
		}
	}

	return nil
}

func (m *meter) getMetricValue(metricFamily *dto.MetricFamily, labels prometheus.Labels) (float64, error) {
	if metricFamily == nil {
		return NilMetricValue, ErrorMetricNotFound
	}

	for _, item := range metricFamily.Metric {
		// compare the number of labels
		if len(item.Label) != len(labels) {
			continue
		}
		// compare the value of labels
		isSame := true
		for _, label := range item.Label {
			if value, exist := labels[label.GetName()]; !exist || value != label.GetValue() {
				isSame = false
				break
			}
		}
		if isSame {
			return m.getValue(metricFamily.GetType(), item), nil
		}
	}

	return NilMetricValue, ErrorMetricNotFound
}

func (m *meter) getValue(metricType dto.MetricType, metric *dto.Metric) float64 {
	switch metricType {
	case dto.MetricType_COUNTER:
		return metric.GetCounter().GetValue()
	case dto.MetricType_GAUGE:
		return metric.GetGauge().GetValue()
	case dto.MetricType_SUMMARY:
		return metric.GetSummary().GetSampleSum()
	case dto.MetricType_UNTYPED:
		return metric.GetUntyped().GetValue()
	case dto.MetricType_HISTOGRAM:
		return metric.GetHistogram().GetSampleSum()
	default:
		clog.Blog.Errorf("unsupport the metric kind: %s", metricType)
	}
	return NilMetricValue
}

func (m *meter) updateGaugeVec() {
	m.lock.Lock()
	defer m.lock.Unlock()

	curMetric := m.getMetricFamily(dto.MetricType_COUNTER, m.name+CurCounterSuffix)
	lastMetric := m.getMetricFamily(dto.MetricType_COUNTER, m.name+LastCounterSuffix)
	if curMetric == nil {
		return
	}
	for _, metricItem := range curMetric.Metric {
		// get curMetric value
		var curValue = metricItem.GetCounter().GetValue()
		// construct label selector
		var labels prometheus.Labels = make(map[string]string)
		for _, labelPair := range metricItem.Label {
			labels[labelPair.GetName()] = labelPair.GetValue()
		}
		// get lastMetric value with same labels
		lastValue, err := m.getMetricValue(lastMetric, labels)
		if err == ErrorMetricNotFound {
			lastValue = 0
		}
		// query per second
		qps := (curValue - lastValue) / float64(m.interval)
		// set gauge metric value
		m.gauge.With(labels).Set(qps)

		// rotate -> copy lastCounter from curCounter after loop
		m.last.With(labels).Add(curValue - lastValue)
	}
}

func (m *meter) meterGC() {
	m.lock.Lock()
	defer m.lock.Unlock()

	curMetric := m.getMetricFamily(dto.MetricType_COUNTER, m.name+CurCounterSuffix)
	if curMetric == nil {
		return
	}

	isOverweight := len(curMetric.Metric) < DefaultMaxMetricCount
	isTimeout := time.Now().Sub(m.lastGCTime).Minutes() < DefaultGCPeriod
	if isOverweight && isTimeout {
		return
	}

	m.cur.Reset()
	m.last.Reset()
	m.gauge.Reset()
	m.lastGCTime = time.Now()
}

func (m *meter) tick() {
	wait := time.After(0)
	for {
		select {
		case <-wait:
			// update gaugevec value
			m.updateGaugeVec()
			//
			m.meterGC()
		}
		wait = time.After(time.Duration(m.interval) * time.Second)
	}
}
