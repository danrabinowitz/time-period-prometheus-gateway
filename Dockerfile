FROM golang:latest as build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /app/app ./cmd/time-period-prometheus-gateway

################################################################################
FROM gcr.io/distroless/base-debian10
COPY --from=build /app/app /
CMD ["/app"]
