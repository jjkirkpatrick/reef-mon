# Reef-Mon 

Reef-Mon is a simple tool to monitor the status of a Reef using Go, InfluxDB and a Raspberry PI.

# Requirments

Go: https://golang.org/
InfluxDB: https://influxdb.com/

## Installation

Clone the repository and run the following command to start the application:

```bash
cd reef-mon/commands
go run . 
```

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
  monitor1:
    name: "Temperature_1"
    type: "GPIO"
    measurement: "reef_temperature"
    field: "temperature"
    interval: 2
  monitor2:
    name: "Temperature_2"
    type: "GPIO"
    measurement: "reef_temperature"
    field: "temperature"
    interval: 20
```

## Contributing
