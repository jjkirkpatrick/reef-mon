package models

type Monitor struct {
	Monitors map[string]MonitorConfig
}

type MonitorConfig struct {
	Name        string
	Type        MonitorType
	Influx      InfluxConfig
	Measurement string
	Field       string
	Interval    int
}

type MonitorType struct {
	Name     string
	Pin      int
	DeviceID string
}

type InfluxConfig struct {
	Bucket string
}
