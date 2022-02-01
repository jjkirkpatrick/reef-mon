package influxdb

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"gopkg.in/yaml.v2"
)

type InfluxDB struct {
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

type DataPoint struct {
	Measurement string
	Tags        map[string]string
	Fields      map[string]interface{}
	Timestamp   time.Time
}

func New() InfluxDB {
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	influx, err := influxConfig(exPath + "/config.yml")

	if err != nil {
		panic(err)
	}

	return influx
}

func (influx *InfluxDB) Write(datapoint DataPoint) {
	client := influxdb2.NewClient(influx.Influx.Host+":"+influx.Influx.Port, influx.Influx.Token)

	writeAPI := client.WriteAPIBlocking(influx.Influx.Org, influx.Influx.Bucket)

	err := writeAPI.WritePoint(context.Background(), influxdb2.NewPoint(datapoint.Measurement, datapoint.Tags, datapoint.Fields, datapoint.Timestamp))

	if err != nil {
		fmt.Println(err)
	}

	client.Close()
}

// NewConfig returns a new decoded Config struct
func influxConfig(configPath string) (InfluxDB, error) {
	// Create config structure
	config := &InfluxDB{}
	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return InfluxDB{}, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return InfluxDB{}, err
	}

	return *config, nil
}
