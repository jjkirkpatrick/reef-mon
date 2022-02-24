package helpers

import (
	"io/ioutil"
	"log"

	"github.com/jjkikrpatrick/reef-mon/system/models"
	"gopkg.in/yaml.v2"
)

func ReadMonitors(configPath string) models.Monitor {

	monitor := models.Monitor{}

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &monitor)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return monitor

}

func ReadDevices(configPath string) models.Device {

	device := models.Device{}

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &device)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return device

}
