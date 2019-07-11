package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type fooCollector struct {
	fooMetric	*prometheus.Desc
	barMetric	*prometheus.Desc
}

func newFooCollector() *fooCollector {
	return &fooCollector{
		fooMetric: prometheus.NewDesc("foo_metric",
			"show whether a foo has occurred in our cluster",
			nil, nil,
		),
		barMetric: prometheus.NewDesc("bar_metric",
			"show whether a bar has occurred in our cluster",
			nil, nil,
		),
	}
}

func (collector *fooCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.fooMetric
	ch <- collector.barMetric
}
func (collector *fooCollector) Collect(ch chan<- prometheus.Metric) {
	var metricValue float64
	if 1 == 1 {
		metricValue = 8888
	}
	
	ch <- prometheus.MustNewConstMetric(collector.fooMetric, prometheus.CounterValue, metricValue)
	ch <- prometheus.MustNewConstMetric(collector.barMetric, prometheus.CounterValue, 10086)
}