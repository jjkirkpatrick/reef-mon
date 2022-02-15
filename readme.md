# Reef-Mon 

Reef-Mon is a simple tool to monitor the status of a Reef using Go, InfluxDB and a Raspberry PI.

# Requirments

Go: https://golang.org/
InfluxDB: https://influxdb.com/

## Installation

Clone the repository and run the following command to start the application:

```bash
make run-dev
```
This will build the application and start it.


To list all temperature sensors run the following:

```bash
cd reef-mon/commands
go run . list-temperature-devices
```


## Configuration

InfluxDB is used to store the data. The following configuration is required:

```YML
influxDB:
  host: "http://localhost"
  port: "8086"
  token: "mzf...kUWg=="
  bucket: "reef-mon"
  org: "reef-mon"
```

The below configuration is an example of how to configure a monitor.

Name: The name of the monitor, this is used to identify the monitor in the database.

Mesurment: The name of the measurement in the database.

Field: The mesurment field name.

Interval: The interval in seconds used to poll for data.

Type: Yaml representation of the monitor type and confiuration  
type:
  name: "gpio_switch"
  pin: 22

```YML
monitors:
  lowWater1:
    name: "WaterLow"
    measurement: "water_low"
    field: "status"
    interval: 10
    type:
      name: "gpio_switch"
      pin: 22
      active_low: true  
    influx:
      bucket: "reef-mon-water-level"
  temperature1:
    name: "temperature 1"
    measurement: "water_temperature"
    field: "temperature"
    interval: 5
    type:
      name: "1w_temperature"
      deviceid: 28-01211358c5ab
    influx:
      bucket: "reef-mon-water-temperature"
```

## Contributing
