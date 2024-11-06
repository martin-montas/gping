package main

import (
        "flag"
        ping "gping/ICMPRequest"
)

var ip *string = flag.String("ip", "127.0.0.1", "IP address to ping")

func main() {
		flag.Parse()
		ping.Run(ip)
}
