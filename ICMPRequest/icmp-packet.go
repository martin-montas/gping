package ICMPRequest

import (
	"fmt"
	"syscall"
	"net"
	"os"

)
var icmpPacket = []byte{
	8, 0, 0, 0, 
	0, 1,
	0, 1,
	72, 101, 108, 108, 111, 
}

func calculateChecksum() uint16 {
	var sum uint32
	for i := 0; i < len(icmpPacket)-1; i += 2 {
		sum += uint32(icmpPacket[i])<<8 | uint32(icmpPacket[i+1])
	}
	if len(icmpPacket)%2 == 1 {
		sum += uint32(icmpPacket[len(icmpPacket)-1]) << 8
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	return uint16(^sum)
}

func sendSingleIP(ip string) {
	pd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP) 

}

func RunProgram(ip string) {
	isValidisCidr4 := isCidrVAlidIpv4(ip)
	isValidisCidr6  :=  isCidrVAlidIpv6(ip)
	if isValidisCidr4 {
		handleCidr4(ip)
	} 
	if isValidisCidr6 {
		handleCidr6(ip)
	} else {
		sendSingleIP(ip)
	}
}

func readCurrentPing(conn net.Conn) {
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

