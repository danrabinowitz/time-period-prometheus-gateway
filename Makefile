GITCOMMIT = $(shell git rev-parse --short HEAD 2>/dev/null)
NAME=time-period-prometheus-gateway
IMAGE_OWNER=danrabinowitz
DOCKER_REGISTRY=docker-registry.djrtechconsulting.com
PROD_TAG=${DOCKER_REGISTRY}/${IMAGE_OWNER}/${NAME}:$(GITCOMMIT)

.PHONY: build
build: time-period-prometheus-gateway

time-period-prometheus-gateway: cmd/time-period-prometheus-gateway/*.go
	go build ./cmd/time-period-prometheus-gateway

.PHONY: run
run: time-period-prometheus-gateway
	./time-period-prometheus-gateway

.PHONY: curl_test
curl_test:
	curl -s 'http://localhost:9130/metrics' | grep mnFoo

.PHONY: docker
docker:
	docker build -t ${PROD_TAG} .

.PHONY: run-docker
run-docker:
	docker run -v "$(shell pwd -P)/config.yml:/config.yml" -p 9130:9130 time-period-prometheus-gateway

.PHONY: publish
publish: docker
	docker push ${PROD_TAG}
