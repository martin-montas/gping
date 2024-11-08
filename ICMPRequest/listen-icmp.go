package ICMPRequest

import (
	"fmt"
	"log"
	"net"
)

func listenForICMP(ip string) { 
	conn, err := net.ListenPacket("ip4:icmp", ip)
	if err != nil {
		fmt.Printf("Error listening: %v\n", err)
		return
	}
	defer conn.Close()  
	for {
		var buf []byte
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Println("Error reading from connection:", err)
			continue  
		}
		log.Printf("message = '%s' "+
		"length = %d " + 
		"source-ip = %s\n", string(buf[:n]), n, addr)
	}
}
