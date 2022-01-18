package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/jasonlvhit/gocron"
)

func main() {

	cfg, err := NewConfig(configPath)

	if err != nil {
		panic("Failed Reading configuration file")
	}

	monitors := ReadMonitors(configPath)

	//for each monitor print out the name and type
	for _, v := range monitors.Monitors {
		gocron.Every(uint64(v.Interval)).Second().Do(cfg.do_monitor, v)
	}

	cfg.write_to_influx(DataPoint{
		Measurement: "reef_temperature",
		Tags:        map[string]string{"location": "reef"},
		Fields:      map[string]interface{}{"temperature": 24.85},
		Timestamp:   time.Now(),
	})

	//gocron.Every(1).Second().Do(cfg.write_to_influx)
	<-gocron.Start()
}

func randFlt(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func (cfg *Config) do_monitor(monitor MonitorConfig) {

	primer := randFlt(20, 50)
	value := randFlt(primer, primer+5)

	datapoint := DataPoint{
		Measurement: monitor.Measurement,
		Tags:        map[string]string{"name": monitor.Name},
		Fields:      map[string]interface{}{monitor.Field: value},
		Timestamp:   time.Now(),
	}

	cfg.write_to_influx(datapoint)
}

func (cfg *Config) write_to_influx(datapoint DataPoint) {
	client := influxdb2.NewClient(cfg.Influx.Host+":"+cfg.Influx.Port, cfg.Influx.Token)

	writeAPI := client.WriteAPIBlocking(cfg.Influx.Org, cfg.Influx.Bucket)

	err := writeAPI.WritePoint(context.Background(), influxdb2.NewPoint(datapoint.Measurement, datapoint.Tags, datapoint.Fields, datapoint.Timestamp))

	if err != nil {
		fmt.Println(err)
	}

	client.Close()
}
