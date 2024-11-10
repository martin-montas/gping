package ICMPRequest

import (
	"fmt"
	"syscall"
	"log"
)

func listenForICMP(sockfd int) {
	var buffer [2048]byte
	for {
		n, _, err := syscall.Recvfrom(sockfd, buffer[:], 0)
		if err != nil {
			log.Fatal("Error receiving packet:", err)
		}
		fmt.Printf("Captured %d bytes: %x\n", n, buffer[:n])
		break
	}

}

