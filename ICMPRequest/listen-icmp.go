package ICMPRequest

import (
	"fmt"
	"syscall"
	"log"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func listenForICMP(sockfd int) {
	var buffer [2048]byte
	n, from, err := syscall.Recvfrom(sockfd, buffer[:], 0)
	if err != nil {
		log.Fatal("Error receiving packet:", err)
	}
	fmt.Printf("Captured %d bytes: %x\n", n, buffer[:n])
	interpretPacket(buffer[:n], n, from)
}

func interpretPacket(buffer []byte, n int, from syscall.Sockaddr) {
	msg, err := icmp.ParseMessage(ipv4.ICMPTypeEcho.Protocol(), buffer[:n])
	if err != nil {
		log.Printf("Error parsing ICMP message: %v", err)
	}
	if msg.Type == ipv4.ICMPTypeEcho {
		fmt.Printf("Received ICMP Echo Request from %v\n", from)
		if echo, ok := msg.Body.(*icmp.Echo); ok {
			fmt.Printf("ID: %d, Seq: %d, Data: %s\n", echo.ID, echo.Seq, string(echo.Data))
		}
	} else {
		fmt.Printf("Received non-echo ICMP message of type %v\n", msg.Type)
	}
}
