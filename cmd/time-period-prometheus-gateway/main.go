package main

import (
	"time-period-prometheus-gateway/internal/config"
	"time-period-prometheus-gateway/internal/metrics"
	"time-period-prometheus-gateway/internal/server"
)

func main() {
	config := config.New()
	metrics.Register(config)
	server.Run(config)
}
