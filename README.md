# ipqr

**Show your current IP address as a QR code in the terminal**

This makes it easy to test a local site on mobile -- because typing IP addresses on mobile sucks

![ipqr demo](/../screenshots/basic_usage.png?raw=true "ipqr Demo")

## Installation
To get started, make sure you have [go](https://golang.org/) installed, then:

```bash
$ go get github.com/jordwest/ipqr
$ ipqr
```

## Usage

```bash
$ ipqr
```

That's it.

## Examples

#### Display help
```bash
$ ipqr --help
```

```
Usage of ipqr:
  -h, --host string       Override host. This will default to the autodetected IP of this device
  -n, --interface int     The number of the interface to display. Use --list to find the interface number (default -1)
  -l, --list              Show a complete list of detected network addresses. By default we'll try to auto detect
  -a, --path string       Specify a path at the end of the URL
  -p, --port int          The port number to append to the end of the host, if any (default -1)
  -r, --protocol string   The protocol to prepend (default "http")
```

#### Use a specific port and path
```bash
$ ipqr --port 8080 --path /blog
```

```
en0: 192.168.1.226 ==> http://192.168.1.226:8080/blog
```
#### Display local interfaces as a list

```bash
$ ipqr --list
```

```
0: lo0 ::1
1: lo0 127.0.0.1
2: lo0 fe80::1
3: en0 fe80::8a23:9ff:fa4c:48c3
4: en0 192.168.1.226
```

#### Select a specific interface number
```bash
$ ipqr -n 1
```

```
lo0: 127.0.0.1 ==> http://127.0.0.1
```
