package exporter

import (
	"errors"
	"net/url"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func PrometheusFetcherTest(u *url.URL) (float64, error) {
	return 1.0, nil
}
func PrometheusFetcherError(u *url.URL) (float64, error) {
	return 0, errors.New("fake error")
}

func Test_NewExporter(t *testing.T) {
	promAPIQueryURL, _ := url.Parse("http://example.com/")
	_, err := New("namespace", "metricName", "queryTemplate", promAPIQueryURL, "period", PrometheusFetcherTest)

	if err != nil {
		t.Errorf("queryFromTemplate returned an unexpected error: %v", err)
	}
}

func Test_Collect(t *testing.T) {
	promAPIQueryURL, _ := url.Parse("http://example.com/")
	e, _ := New("namespace", "metricName", "queryTemplate", promAPIQueryURL, "current_calendar_month", PrometheusFetcherTest)
	ch := make(chan<- prometheus.Metric)

	go e.Collect(ch)
}

func Test_CollectError(t *testing.T) {
	promAPIQueryURL, _ := url.Parse("http://example.com/")
	e, _ := New("namespace", "metricName", "queryTemplate", promAPIQueryURL, "current_calendar_month", PrometheusFetcherError)
	ch := make(chan<- prometheus.Metric)

	e.Collect(ch)
}

func Test_CollectBadPeriod(t *testing.T) {
	promAPIQueryURL, _ := url.Parse("http://example.com/")
	e, _ := New("namespace", "metricName", "queryTemplate", promAPIQueryURL, "period", PrometheusFetcherTest)
	ch := make(chan<- prometheus.Metric)

	e.Collect(ch)
}
