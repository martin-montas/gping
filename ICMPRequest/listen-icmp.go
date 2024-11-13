package ICMPRequest

import (
	"fmt"
	"syscall"
	"log"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func listenForICMP(sockfd int, ip string) {
	var buffer [2048]byte
	n, _, err := syscall.Recvfrom(sockfd, buffer[:], 0)
	if err != nil {
		log.Fatal("Error receiving packet:", err)
	}
	interpretPacket(buffer, n, ip)
	fmt.Printf("Captured %d bytes: %x\n", n, buffer[:n])
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
