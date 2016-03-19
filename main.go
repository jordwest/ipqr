package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"log"
	"os"
	"strconv"
)

var invert bool
var list bool
var selectedInterface int
var protocol string
var host string
var path string
var config *Config

func init() {
	defaultInvert := false
	defaultInterface := -1

	defaultPort := -1
	defaultHost := ""
	defaultPath := ""
	defaultProtocol := "http"

	if configExists("./.ipqr") {
		config, err := loadConfig(filename)
	} else {
		config := &Config{}
	}

	envPort := os.Getenv("PORT")
	if envPort != "" {
		port, err := strconv.Atoi(envPort)
		if err != nil {
			log.Printf("Warning: Invalid PORT environment variable set: %s", envPort)
		} else {
			config.Port = port
		}
	}

	flag.BoolVarP(&invert, "invert", "i", defaultInvert, "Use if your terminal has a light background")
	flag.BoolVarP(&list, "list", "l", false, "Show a complete list of detected network addresses. By default we'll try to auto detect")
	flag.BoolVarP(&save, "save", "s", false, "Save the passed URL options to a configuration file in the cwd")
	flag.IntVarP(&selectedInterface, "interface", "n", defaultInterface, "The number of the interface to display. Use --list to find the interface number")
	flag.IntVarP(&config.Port, "port", "p", defaultPort, "The port number to append to the end of the host, if any")
	flag.StringVarP(&config.Protocol, "protocol", "r", defaultProtocol, "The protocol to prepend")
	flag.StringVarP(&config.Host, "host", "h", defaultHost, "Override host. This will default to the autodetected IP of this device")
	flag.StringVarP(&config.Path, "path", "a", defaultPath, "Specify a path at the end of the URL")

	flag.Parse()

}

func printList(options []Option) {
	for i, option := range options {
		fmt.Printf("%d: %s %s\n", i, option.iface, option.ip.String())
	}
}

func printOption(o Option) {
	url := o.MakeURL(protocol, port, path)
	err := PrintQRCode(url)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	fmt.Printf("%s: %s ==> %s\n", o.iface, o.ip.String(), url)
}

func main() {
	if len(host) > 0 {
		url := MakeURL(protocol, host, port, path)
		err := PrintQRCode(url)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		fmt.Printf("%s\n", url)
		return
	}

	options, err := DetectOptions()
	if err != nil {
		log.Fatal(err)
	}

	if list {
		printList(options)
		return
	}

	if selectedInterface != -1 {
		option, err := options.Get(selectedInterface)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		printOption(option)
		return
	}

	// Autodetect
	option, err := options.Autodetect()
	if err != nil {
		fmt.Printf("Could not autodetect an interface. Try re-running with the --list flag.")
	}
	printOption(option)
}
