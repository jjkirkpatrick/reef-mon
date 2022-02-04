package temperature

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/jjkikrpatrick/reef-mon/system/drivers/influxdb"
	"github.com/jjkikrpatrick/reef-mon/system/models"
)

var ErrReadSensor = errors.New("failed to read sensor temperature")

func Get(monitorConfig models.MonitorConfig) {
	influx := influxdb.New()
	sensors, err := sensors()

	if err != nil {
		fmt.Printf("Error reading sensors: %s\n", err)
		return
	}

	for _, sensor := range sensors {
		if sensor == monitorConfig.Type.DeviceID {
			temperature, err := temperature(sensor)

			if err != nil {
				fmt.Println(err)
			}

			datapoint := influxdb.DataPoint{
				Measurement: monitorConfig.Measurement,
				Tags:        map[string]string{"name": monitorConfig.Name},
				Fields:      map[string]interface{}{monitorConfig.Field: temperature},
				Timestamp:   time.Now(),
			}
			influx.Write(datapoint)
		}
	}
}

// Sensors get all connected sensor IDs as array
func sensors() ([]string, error) {
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
func temperature(sensor string) (float64, error) {
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

func ListTemperatureDevices() [][]string {
	sensors, err := sensors()

	results := [][]string{}

	if err != nil {
		fmt.Printf("Error reading sensors: %s\n", err)
		return nil
	}

	for _, sensor := range sensors {
		temperature, err := temperature(sensor)
		if err != nil {
			fmt.Println(err)
		}
		results = append(results, []string{sensor, fmt.Sprintf("%.2f", temperature)})
	}
	return results
}
