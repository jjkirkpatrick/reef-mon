package system

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jasonlvhit/gocron"
	"github.com/jjkikrpatrick/reef-mon/system/internal/helpers"
	"github.com/jjkikrpatrick/reef-mon/system/modules/gpio"
	"github.com/jjkikrpatrick/reef-mon/system/modules/temperature"
	"github.com/jjkikrpatrick/reef-mon/system/modules/camera"
)

func Start() chan bool {

	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	monitors := helpers.ReadMonitors(exPath + "/config.yml")
	devices := helpers.ReadDevices(exPath + "/config.yml")

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

	for _, v := range devices.Devices {
		if v.Type.Name == "camera" {
			camera.Stream( v.Type.Http_addr, v.Type.Device_id, v.Type.Uri,  false)
		}
	}



	return nil
}
