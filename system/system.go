package system

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
	"github.com/jjkikrpatrick/reef-mon/system/internal/helpers"
	"github.com/jjkikrpatrick/reef-mon/system/modules/gpio"
	"github.com/jjkikrpatrick/reef-mon/system/modules/temperature"
)

func Start() chan bool {

	const configPath = "../config/config.yml"

	monitors := helpers.ReadMonitors(configPath)

	//for each monitor print out the name and type
	for _, v := range monitors.Monitors {

		if v.Type.Name == "gpio_switch" {
			gocron.Every(uint64(v.Interval)).Second().Do(gpio.GetSwitch, v)
		} else if v.Type.Name == "1w_temperature" {
			gocron.Every(uint64(v.Interval)).Second().Do(temperature.Get, v)
		} else {
			fmt.Printf("Unknown monitor type: %s for monitor: %s \n", v.Type, v.Name)
		}

	}

	gocron.Start()

	return nil
}
