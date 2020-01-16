GITCOMMIT = $(shell git rev-parse --short HEAD 2>/dev/null)
NAME=time-period-prometheus-gateway
IMAGE_OWNER=danrabinowitz
DOCKER_REGISTRY=docker-registry.djrtechconsulting.com
PROD_TAG=${DOCKER_REGISTRY}/${IMAGE_OWNER}/${NAME}:$(GITCOMMIT)

.PHONY: build
build:
	go build ./cmd/time-period-prometheus-gateway

.PHONY: run
run: build
	./time-period-prometheus-gateway

.PHONY: curl_test
curl_test:
	curl -s 'http://localhost:9130/metrics' | grep unifi_wan_receive_bytes_total

.PHONY: docker
docker: test
	docker build -t ${PROD_TAG} .

.PHONY: run-docker
run-docker:
	docker run -v "$(shell pwd -P)/config.yml:/config.yml" -p 9130:9130 time-period-prometheus-gateway

.PHONY: publish
publish: docker
	docker push ${PROD_TAG}

.PHONY: test
test:
	go test ./...

.PHONY: coverage
coverage:
	go test ./... -coverprofile cover.out
	go tool cover -func cover.out
	# go tool cover -html cover.out
