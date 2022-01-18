package main

import "time"

type Config struct {
	Influx struct {
		// Host is the local machine IP Address to bind the HTTP Server to
		Host string `yaml:"host"`

		// Port is the local machine TCP Port to bind the HTTP Server to
		Port   string `yaml:"port"`
		Token  string `yaml:"token"`
		Org    string `yaml:"org"`
		Bucket string `yaml:"bucket"`
	} `yaml:"influxDB"`
}

type Monitor struct {
	Monitors map[string]MonitorConfig
}

type MonitorConfig struct {
	Name        string
	Type        string
	Measurement string
	Field       string
	Interval    int
}

type DataPoint struct {
	Measurement string
	Tags        map[string]string
	Fields      map[string]interface{}
	Timestamp   time.Time
}
