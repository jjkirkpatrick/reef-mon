package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"github.com/jjkikrpatrick/reef-mon/system"
	"github.com/jjkikrpatrick/reef-mon/system/modules/temperature"
)

var Version string

func main() {

	version := flag.Bool("version", false, "Print version information")
	flag.Usage = func() {
		text := `reef-mon is a tool for monitoring the parameters of a reef.`
		fmt.Println(strings.TrimSpace(text))
	}
	flag.Parse()
	if *version {
		fmt.Println(Version)
		return
	}

	var v string
	args := []string{}
	if len(os.Args) > 1 {
		v = os.Args[1]
	}
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
	switch v {
	case "", "start":
		fmt.Println("Starting reef-Mon")
		<-system.Start()
	case "list-temperature-devices":
		fmt.Println("list-temperature-devices", args)
		results := temperature.ListTemperatureDevices()
		for _, result := range results {
			//TODO: Properly cast result[1] to float64
			fmt.Printf("%s : %s°c\n", result[0], result[1])
		}
	default:
		fmt.Println("Unknown command: '", v, "'")
		os.Exit(1)
	}
}