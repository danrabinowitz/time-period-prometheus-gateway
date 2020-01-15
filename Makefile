build: time-period-prometheus-gateway

time-period-prometheus-gateway: cmd/time-period-prometheus-gateway/exporter.go cmd/time-period-prometheus-gateway/main.go cmd/time-period-prometheus-gateway/prometheus_fetcher.go cmd/time-period-prometheus-gateway/period_processor.go
	go build ./cmd/time-period-prometheus-gateway

run: time-period-prometheus-gateway
	./time-period-prometheus-gateway

.PHONY: curl_test
curl_test:
	curl -s 'http://localhost:9130/metrics' | grep mnFoo
