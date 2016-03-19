package main

import (
	"fmt"
	"github.com/boombuler/barcode/qr"
	"log"
	"net"
)

func white(times int) {
	for i := 0; i < times; i++ {
		fmt.Printf("\x1b[7m \x1b[0m")
	}
}
func black(times int) {
	for i := 0; i < times; i++ {
		fmt.Printf(" ")
	}
}

func printQRCode(data string) {
	qrcode, err := qr.Encode(data, qr.Q, qr.Auto)
	if err != nil {
		log.Fatalf("Error generating QR code %s\n", err)
	}

	border := 1
	charWidth := 2

	rect := qrcode.Bounds()
	fmt.Printf("%dx%d\n", rect.Dx(), rect.Dy())
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

}

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatalf("Error querying interfaces: %s", err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Printf("Error querying interface addresses: %s", err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip.IsGlobalUnicast() {
				printQRCode(fmt.Sprintf("http://%s:1313", ip))
				fmt.Printf("%s: %s\n", i.Name, ip)
			}
		}
	}
}
