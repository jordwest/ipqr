package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"log"
)

var list bool
var selectedInterface int
var port int
var protocol string
var host string
var path string

func init() {
	flag.BoolVarP(&list, "list", "l", false, "Show a complete list of detected network addresses. By default we'll try to auto detect")
	flag.IntVarP(&selectedInterface, "interface", "n", -1, "The number of the interface to display. Use --list to find the interface number")
	flag.IntVarP(&port, "port", "p", -1, "The port number to append to the end of the host, if any")
	flag.StringVarP(&protocol, "protocol", "r", "http", "The protocol to prepend")
	flag.StringVarP(&host, "host", "h", "", "Override host. This will default to the autodetected IP of this device")
	flag.StringVarP(&path, "path", "a", "", "Specify a path at the end of the URL")
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
	flag.Parse()

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
