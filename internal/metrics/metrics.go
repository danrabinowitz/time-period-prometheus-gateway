package metrics

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"

	"time-period-prometheus-gateway/internal/config"
	"time-period-prometheus-gateway/internal/exporter"
	"time-period-prometheus-gateway/internal/prometheusclient"
)

// Register registers metrics
func Register(config config.Config) {
	for _, m := range config.Metrics {
		metricName := m.Name
		if metricName == "" {
			log.Panicf("metric_name is required")
		}
		queryTemplate := m.QueryTemplate
		// log.Printf("queryTemplate=%q", queryTemplate)
		if queryTemplate == "" {
			log.Panicf("query_template is required")
		}
		period := m.Period
		if period == "" {
			log.Panicf("period is required")
		}

		e, err := exporter.New(config.Namespace, metricName, queryTemplate, config.PromAPIQueryURL, period, prometheusclient.PrometheusFetcher)
		if err != nil {
			log.Panicf("failed to create exporter: %v", err)
		}

		prometheus.MustRegister(e)
	}
}
