package main

import (
	"log"
	"net/url"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

// An exporter is a Prometheus exporter
type exporter struct {
	mu              sync.Mutex
	namespace       string
	metricName      string
	queryTemplate   string
	promAPIQueryURL url.URL
	period          string
}

// New creates a new Exporter
func newExporter(namespace string, metricName string, queryTemplate string, promAPIQueryURL *url.URL, period string) (*exporter, error) {

	e := &exporter{
		namespace:       namespace,
		metricName:      metricName,
		queryTemplate:   queryTemplate,
		promAPIQueryURL: *promAPIQueryURL,
		period:          period,
	}

	return e, nil
}

// Collect sends the collected metrics from each of the collectors to
// prometheus. Collect could be called several times concurrently
// and thus its run is protected by a single mutex.
func (e *exporter) Collect(ch chan<- prometheus.Metric) {
	e.mu.Lock()
	defer e.mu.Unlock()

	baseLabels := []string{}
	labels := []string{}

	promDesc := prometheus.NewDesc(
		prometheus.BuildFQName(e.namespace, e.period, e.metricName),
		e.metricName,
		baseLabels,
		nil,
	)

	v, err := e.value()
	if err != nil {
		log.Printf("Error fetching value %q", err)
		return
	}

	ch <- prometheus.MustNewConstMetric(
		promDesc,
		prometheus.GaugeValue,
		v,
		labels...,
	)
}

func (e *exporter) value() (float64, error) {
	queryParam, err := queryFromTemplate(e.queryTemplate, e.period)
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}

	u := e.promAPIQueryURL
	// Query params
	params := url.Values{}
	params.Add("query", queryParam)
	u.RawQuery = params.Encode()

	// log.Printf("url=%q", u.String())

	return prometheusFetcher(&u)
}

// Describe sends all the descriptors of the collectors included to
// the provided channel.
func (e *exporter) Describe(ch chan<- *prometheus.Desc) {
	e.mu.Lock()
	defer e.mu.Unlock()

	baseLabels := []string{}

	promDesc := prometheus.NewDesc(
		prometheus.BuildFQName(e.namespace, e.period, e.metricName),
		e.metricName,
		baseLabels,
		nil,
	)

	ch <- promDesc
}
