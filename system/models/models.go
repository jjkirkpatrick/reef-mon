package models

//Monitors

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
	Name       string
	Pin        int
	DeviceID   string
	Active_low bool
}

type InfluxConfig struct {
	Bucket string
}

//devices

type Device struct {
	Devices map[string]DeviceConfig
}

type DeviceConfig struct {
	Name        string
	Type        DeviceType
}

type DeviceType struct {
	Name       	string
	Device_id	string
	Http_addr   string
	Uri 	  	string
}