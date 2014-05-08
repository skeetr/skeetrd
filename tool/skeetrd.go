package main

import (
	"flag"
	"fmt"
	. "skeetrd"
	"time"
)

type Options struct {
	configFile string
	debug      bool
	help       bool
}

var version string
var options Options

func init() {
	flag.StringVar(&options.configFile, "config", "/etc/skeetrd.conf", "config filename")
	flag.BoolVar(&options.debug, "debug", false, "raise log level to debug")
	flag.BoolVar(&options.help, "help", false, "display this help")

	flag.Usage = help
}

func main() {
	flag.Parse()

	switch {
	case options.help:
		help()
	default:
		run()
	}
}

func help() {
	fmt.Printf("\033[1mskeetrd (%s)\033[0m\n", version)
	fmt.Printf("Low footprint collector and parser for events and logs\n")
	fmt.Printf("Máximo Cuadros Ortiz <mcuadros@gmail.com>\n\n")

	fmt.Printf("Usage:\n")
	flag.PrintDefaults()
}

func run() {
	server := NewServer(&ServerConfig{
		Port: 1234,
		Host: "localhost",
	})

	server.Start()

	time.Sleep(60 * time.Second)
}
