package main

import (
        "flag"
        gping "gping/ICMPRequest"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "IP address to ping")

	flag.Parse()
	gping.RunProgram(*ip)
}
