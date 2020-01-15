package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/prometheus/client_golang/prometheus"
	// "github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/yaml.v2"
)

type metric struct {
	Name          string `yaml:"name"`
	QueryTemplate string `yaml:"query_template"`
	Period        string `yaml:"period"`
}

type config struct {
	Listen             map[string]string `yaml:"listen"`
	Metrics            []metric          `yaml:"metrics"`
	Namespace          string            `yaml:"namespace"`
	Subsystem          string            `yaml:"subsystem"`
	PromAPIQueryString string            `yaml:"prometheus_api_query_url"`
}

func main() {
	configFile := flag.String("config.file", "", "Relative path to config file yaml")
	if *configFile == "" {
		*configFile = "config.yml"
	}

	flag.Parse()

	var config config
	source, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("failed to read config file %q: %v", *configFile, err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		log.Fatalf("failed to read YAML from config file %q: %v", *configFile, err)
	}

	listenAddr := config.Listen["address"]
	metricsPath := config.Listen["metricspath"]

	if listenAddr == "" {
		listenAddr = ":9130"
	}
	if metricsPath == "" {
		metricsPath = "/metrics"
	}

	namespace := config.Namespace
	subsystem := config.Subsystem

	// promAPIQueryString := "http://localhost:9090/api/v1/query"
	promAPIQueryString := config.PromAPIQueryString
	if promAPIQueryString == "" {
		log.Fatalf("prometheus_api_query_url is required")
	}

	promAPIQueryURL, err := url.Parse(promAPIQueryString)
	if err != nil {
		log.Fatalln(err)
	}

	for _, m := range config.Metrics {
		metricName := m.Name
		if metricName == "" {
			log.Fatalf("metric_name is required")
		}
		queryTemplate := m.QueryTemplate
		// log.Printf("queryTemplate=%q", queryTemplate)
		if queryTemplate == "" {
			log.Fatalf("query_template is required")
		}
		period := m.Period
		if period == "" {
			log.Fatalf("period is required")
		}

		e, err := newExporter(namespace, subsystem, metricName, queryTemplate, promAPIQueryURL, period)
		if err != nil {
			log.Fatalf("failed to create exporter: %v", err)
		}

		prometheus.MustRegister(e)
	}

	http.Handle(metricsPath, promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, metricsPath, http.StatusMovedPermanently)
	})

	log.Printf("Starting Time Period Prometheus Gateway on %q", listenAddr)

	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatalf("Exited Time Period Prometheus Gateway: %s", err)
	}
}
