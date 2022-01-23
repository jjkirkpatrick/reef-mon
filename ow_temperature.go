package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var ErrReadSensor = errors.New("failed to read sensor temperature")

func (cfg *Config) temperature(monitor MonitorConfig) {

	sensors, err := Sensors()

	if err != nil {
		fmt.Printf("Error reading sensors: %s\n", err)
		return
	}

	for _, sensor := range sensors {
		if sensor == monitor.Type.DeviceID {
			temperature, err := Temperature(sensor)

			if err != nil {
				fmt.Println(err)
			}

			datapoint := DataPoint{
				Measurement: monitor.Measurement,
				Tags:        map[string]string{"name": monitor.Name},
				Fields:      map[string]interface{}{monitor.Field: temperature},
				Timestamp:   time.Now(),
			}
			cfg.write_to_influx(datapoint)
		}
	}
}

// Sensors get all connected sensor IDs as array
func Sensors() ([]string, error) {
	data, err := ioutil.ReadFile("/sys/bus/w1/devices/w1_bus_master1/w1_master_slaves")
	if err != nil {
		return nil, err
	}

	sensors := strings.Split(string(data), "\n")
	if len(sensors) > 0 {
		sensors = sensors[:len(sensors)-1]
	}

	return sensors, nil
}

// Temperature get the temperature of a given sensor
func Temperature(sensor string) (float64, error) {
	data, err := ioutil.ReadFile("/sys/bus/w1/devices/" + sensor + "/w1_slave")
	if err != nil {
		return 0.0, ErrReadSensor
	}

	raw := string(data)

	if !strings.Contains(raw, " YES") {
		return 0.0, ErrReadSensor
	}

	i := strings.LastIndex(raw, "t=")
	if i == -1 {
		return 0.0, ErrReadSensor
	}

	c, err := strconv.ParseFloat(raw[i+2:len(raw)-1], 64)
	if err != nil {
		return 0.0, ErrReadSensor
	}

	return c / 1000.0, nil
}
