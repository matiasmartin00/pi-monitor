package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

var Config config

type config struct {
	Metrics metricsConfig `yaml:"metrics"`
}

type metricsConfig struct {
	Cpu       metricConfig `yaml:"cpu"`
	Disk      metricConfig `yaml:"disk"`
	Memory    metricConfig `yaml:"memory"`
	Host      metricConfig `yaml:"host"`
	Speedtest metricConfig `yaml:"speedtest"`
}

type metricConfig struct {
	Enabled  bool           `yaml:"enabled"`
	Interval *time.Duration `yaml:"interval"`
}

func Load() {
	log.Println("Loading configuration")
	data, err := os.ReadFile("configuration.yml")

	if err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
	}

	err = yaml.Unmarshal(data, &Config)

	if err != nil {
		log.Fatalf("Error parsing configuration file: %v", err)
	}

	log.Printf("Configuration loaded: %v", Config)

}
