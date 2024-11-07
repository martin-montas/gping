package main

import (
        "flag"
        ping "gping/ICMPRequest"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "IP address to ping")
	flag.Parse()
	ping.Run(*ip)
}
