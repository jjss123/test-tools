package handler

import (
	"testTools/src/utils/metric"
)

var (
	MatrixRequestRate = metric.NewMeterVec(metric.MeterInterval, metric.MeterOpts{
		Name: "jvessel_matrix_request_rate",
		Help: "QPS of matrix request",
	},
		[]string{metric.MetricLabelUrl},
	)
)
