package ICMPRequest

import (
	"fmt"
	"encoding/binary"
	"log"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func listenForICMP(icmpReply []byte) {
	if icmpReply[0] == 0 { // ICMP Echo Reply
		fmt.Printf("Got echo reply!\n")
		fmt.Printf("Type: %d\n", icmpReply[0])
		fmt.Printf("Code: %d\n", icmpReply[1])
		fmt.Printf("ID: %d\n", binary.BigEndian.Uint16(icmpReply[4:6]))
		fmt.Printf("Sequence: %d\n", binary.BigEndian.Uint16(icmpReply[6:8]))
		fmt.Printf("Data: %s\n", string(icmpReply[8:]))
	}
}

func interpretPacket(buffer [2048]byte, n int, ip string) {
	msg, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), buffer[:n])
	if err != nil {
		log.Printf("Error parsing ICMP message: %v", err)
		return
	}
	fmt.Printf("Source IP: %s\n", ip)
	fmt.Printf("Packet Type: %d\n", msg.Type)
	fmt.Printf("Packet Code: %d\n", msg.Code)
	if msg.Type == ipv4.ICMPTypeEchoReply {
		if echo, ok := msg.Body.(*icmp.Echo); ok {
			fmt.Printf("Echo Data: %s\n", string(echo.Data))
		} else {
			log.Println("Error: Unexpected body type")
		}
	}
}
