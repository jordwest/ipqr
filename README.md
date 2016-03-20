# ipqr

**Show your current IP address as a QR code in the terminal**

This makes it easy to test a local site on mobile -- because typing IP addresses on mobile sucks

![ipqr demo](/../screenshots/basic_usage.png?raw=true "ipqr Demo")

## Download

### Mac

Using homebrew:

```bash
brew install jordwest/tools/ipqr
```

Alternatively, download the latest `darwin_amd64.tar.gz` package from the [Releases page](https://github.com/jordwest/ipqr/releases/)

### Windows

Download the latest `windows_amd64.zip` or `windows_386.zip` from the [Releases page](https://github.com/jordwest/ipqr/releases/).

### Linux (including Raspberry Pi!)

Download the latest `linux_*.tar.gz` package from the [Releases page](https://github.com/jordwest/ipqr/releases/).

For x86/x86_64 processors, use `linux_386.tar.gz` or `linux_amd64.tar.gz`.

For Raspberry Pi, use `linux_arm.tar.gz`. Note that although a 64-bit binary is available, the default OS on the Raspberry Pi 3 (which has a 64-bit processor) is still 32-bit.

## Installation from source
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
  -v, --version           Show the version number and exit
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

#### Customize the host (skips IP detection)

```bash
$ ipqr --host google.com --path /search?q=github+ipqr
```

```
http://google.com/search?q=github+ipqr
```
