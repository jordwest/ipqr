package main

import (
	"fmt"
	"github.com/boombuler/barcode/qr"
	"github.com/shiena/ansicolor"
	"io"
	"net"
	"os"
)

type Option struct {
	iface string
	ip    net.IP
}
type Options []Option

var ansiWriter io.Writer

func init() {
	// For Windows support
	ansiWriter = ansicolor.NewAnsiColorWriter(os.Stdout)
}

func pixel(times int, white bool) {
	for i := 0; i < times; i++ {
		if white {
			fmt.Fprintf(ansiWriter, "\x1b[47m \x1b[0m")
		} else {
			fmt.Fprintf(ansiWriter, "\x1b[40m \x1b[0m")
		}
	}
}

func white(times int) {
	pixel(times, true)
}
func black(times int) {
	pixel(times, false)
}

// Autodetect returns the first global unicast address, or an error if none was found
func (o Options) Autodetect() (Option, error) {
	for _, option := range o {
		if option.ip.IsGlobalUnicast() {
			return option, nil
		}
	}

	return Option{}, fmt.Errorf("Could not autodetect a useful address")
}

// Get selects a specific interface by index, in the order detected
func (o Options) Get(n int) (Option, error) {
	if n < 0 || n > len(o) {
		return Option{}, fmt.Errorf("Interface number out of range")
	}
	return o[n], nil
}

func DetectOptions() (Options, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("Querying interfaces: %s", err)
	}

	var options []Option
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, fmt.Errorf("Querying device addresses: %s", err)
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPAddr:
			case *net.IPNet:
				options = append(options, Option{iface: i.Name, ip: v.IP})
			}
		}
	}

	return options, nil
}

func (o Option) MakeURL(protocol string, port int, path string) string {
	return MakeURL(protocol, o.ip.String(), port, path)
}

func MakeURL(protocol string, host string, port int, path string) string {
	url := fmt.Sprintf("%s://%s", protocol, host)
	if port > 0 {
		url = fmt.Sprintf("%s:%d", url, port)
	}
	if len(path) > 0 {
		if path[0] != '/' {
			url = fmt.Sprintf("%s/", url)
		}

		url = fmt.Sprintf("%s%s", url, path)
	}

	return url
}

func PrintQRCode(data string) error {
	qrcode, err := qr.Encode(data, qr.Q, qr.Auto)
	if err != nil {
		return fmt.Errorf("Generating QR code:", err)
	}

	border := 1
	charWidth := 2

	rect := qrcode.Bounds()
	white(charWidth * (rect.Dx() + border*2))
	fmt.Printf("\n")
	for y := 0; y < rect.Dy(); y++ {
		white(charWidth * border)
		for x := 0; x < rect.Dx(); x++ {

			color := qrcode.At(x, y)
			r, _, _, _ := color.RGBA()
			if r > 128 {
				white(charWidth)
			} else {
				black(charWidth)
			}
		}
		white(charWidth * border)
		fmt.Printf("\n")
	}
	white(charWidth * (rect.Dx() + border*2))
	fmt.Printf("\n")
	return nil
}
