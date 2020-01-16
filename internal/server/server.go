package server

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"time-period-prometheus-gateway/internal/config"
)

// Run runs the server
func Run(config config.Config) {
	http.Handle(config.Listen.MetricsPath, promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, config.Listen.MetricsPath, http.StatusMovedPermanently)
	})

	log.Printf("Starting Time Period Prometheus Gateway on %q", config.Listen.Address)

	if err := http.ListenAndServe(config.Listen.Address, nil); err != nil {
		log.Fatalf("Exited Time Period Prometheus Gateway: %s", err)
	}
}
