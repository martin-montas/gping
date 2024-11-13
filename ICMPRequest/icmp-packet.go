package ICMPRequest

import (
	"fmt"
	"strings"
	"strconv"
	"log"
	"syscall"
)

func checksum(data []byte) uint16 {
	var sum uint32
	for i := 0; i < len(data)-1; i += 2 {
		sum += uint32(data[i])<<8 + uint32(data[i+1])
	}
	if len(data)%2 != 0 {
		sum += uint32(data[len(data)-1]) << 8
	}
	for (sum >> 16) > 0 {
		sum = (sum & 0xFFFF) + (sum >> 16)
	}
	return ^uint16(sum)
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

func handlePacket() []byte {
	var packet = []byte{
		8, 0, 0, 0, 
		0, 1,
		0, 1,
		72, 101, 108, 108, 111, 
	}
	cs := checksum(packet)
	packet[2] = byte(cs >> 8)   
	packet[3] = byte(cs & 0xFF) 
	return packet
}

func stringToByte(ip string) [4]byte {
	strResult := strings.Split(ip, ".")
	var ipByte [4]byte
	for index, value := range strResult {
		intResult, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("Error converting string to int: %v", err)
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
	addr := syscall.SockaddrInet4{
		Addr: stringToByte(ip), 
	}
	packet := handlePacket()
	return sock, packet, addr
}

func sendPacket(ip string) {
	sock, packet, addr := createRawSocket(ip)
	defer syscall.Close(sock)
	err := syscall.Sendto(sock, packet, 0, &addr)
	if err != nil {
		log.Fatalf("Failed to send packet: %v", err)
	}
		listenForICMP(sock)
}
