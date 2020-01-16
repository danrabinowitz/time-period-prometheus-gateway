package config

import (
	"flag"
	"io/ioutil"
	"log"
	"net/url"

	"gopkg.in/yaml.v2"
)

// Config is the config for the command
type Config struct {
	Listen          listen
	Metrics         []metric
	Namespace       string
	PromAPIQueryURL *url.URL
}

// New returns a config which provides the config for the command
func New() Config {
	fileName := flag.String("config.file", "", "Relative path to config file yaml")
	flag.Parse()

	return newConfig(*fileName)
}

func newConfig(fileName string) Config {
	if fileName == "" {
		fileName = "config.yml"
	}

	source, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed to read config file %q: %v", fileName, err)
	}

	var configFile configFile
	err = yaml.UnmarshalStrict(source, &configFile)
	if err != nil {
		log.Panicf("failed to read YAML from config file %q: %v", fileName, err)
	}

	var config Config
	config.Listen = configFile.Listen
	config.Metrics = configFile.Metrics
	config.Namespace = configFile.Namespace

	if config.Listen.Address == "" {
		config.Listen.Address = ":9130"
	}
	if config.Listen.MetricsPath == "" {
		config.Listen.MetricsPath = "/metrics"
	}
	if configFile.PromAPIQueryString == "" {
		log.Panicf("prometheus_api_query_url is required")
	}

	config.PromAPIQueryURL, err = url.Parse(configFile.PromAPIQueryString)
	if err != nil {
		log.Panicf("failed to parse url %q: %v", configFile.PromAPIQueryString, err)
	}

	return config
}

type configFile struct {
	Listen             listen   `yaml:"listen"`
	Metrics            []metric `yaml:"metrics"`
	Namespace          string   `yaml:"namespace"`
	PromAPIQueryString string   `yaml:"prometheus_api_query_url"`
}

type listen struct {
	Address     string `yaml:"address"`
	MetricsPath string `yaml:"metricspath"`
}

type metric struct {
	Name          string `yaml:"name"`
	QueryTemplate string `yaml:"query_template"`
	Period        string `yaml:"period"`
}
