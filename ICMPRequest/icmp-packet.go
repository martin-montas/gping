package ICMPRequest

import (
	"fmt"
	"github.com/google/gopacket"
	"strings"
	"strconv"
	"log"
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


func RunProgram(ip string) {
	isValidisCidr4 := isCidrVAlidIpv4(ip)
	isValidisCidr6  :=  isCidrVAlidIpv6(ip)
	if isValidisCidr4 {
		handleCidr4(ip)
	} 
	if isValidisCidr6 {
		handleCidr6(ip)
	} else {
		sendPacket(ip)
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

func stringToByte(ip string) [4]byte {
	strResult := strings.Split(ip, ".")
	var ipByte [4]byte
	for index, value := range strResult {
		intResult, err := strconv.Atoi(value)
		if err != nil {
			fmt.Errorf("Error converting string to int: %v", err)
		}
		ipByte[index] = byte(intResult)

	}
	return ipByte
}

func createRawSocket(ip string) (int ,[]byte, syscall.SockaddrInet4) {
	sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		log.Fatalf("Failed to create raw socket: %v", err)
	}
	defer syscall.Close(sock)
	addr := syscall.SockaddrInet4{
		Port: 0,
		// Ip goes here
		Addr: stringToByte(ip), 
	}
	packet := []byte{
		8, 0,  // Type and Code for ICMP Echo Request
		0, 0,  // Checksum (simplified; not valid in real use without proper calculation)
		0, 1,  // Identifier
		0, 1,  // Sequence number
		'H', 'e', 'l', 'l', 'o', '!', 
	}
	return sock, packet, addr
}

func sendPacket(ip string) {
	sock, packet, addr := createRawSocket(ip)
	err := syscall.Sendto(sock, packet, 0, &addr)
	if err != nil {
		log.Fatalf("Failed to send packet: %v", err)
	}
	fmt.Println("Packet sent successfully!")
}
