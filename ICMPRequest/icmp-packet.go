package ICMPRequest

import (
	"fmt"
	"net"
	"os"
)

func sendIP(ip string) {
	conn, err := net.Dial("ip4:icmp", ip)
	if err != nil {
		fmt.Println("Error dialing:", err)
		os.Exit(1)
	}
	defer conn.Close()
	readPing(conn)
}

func Run(ip string) {
	// Check if is a notation IP range:
	isValidisCidr4 := isCidrVAlidIpv4(ip)
	isValidisCidr6  :=  isCidrVAlidIpv6(ip)

	if isValidisCidr4 {
		handleCidr4(ip)
	} 

	if isValidisCidr6 {
		handleCidr6(ip)

	} else {
		sendIP(ip)
	}
}

func readPing(conn net.Conn) {
	// Read the response
	buf := make([]byte, 1500)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		os.Exit(1)
	}
	fmt.Printf("Received %d bytes: %v\n", n, buf[:n])
}

func incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
