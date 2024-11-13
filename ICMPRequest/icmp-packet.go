package ICMPRequest

import (
	"encoding/binary"
	"net"
	"strconv"
	"syscall"
	"strings"
	"fmt"
	"os"
)


// ICMP packet structure
type icmpPacket struct {
	Type     uint8
	Code     uint8
	Checksum uint16
	ID       uint16
	Sequence uint16
	Data     []byte
}

func calculateChecksum(data []byte) uint16 {
	var sum uint32
	for i := 0; i < len(data)-1; i += 2 {
		sum += uint32(binary.BigEndian.Uint16(data[i:]))
	}
	if len(data)%2 == 1 {
		sum += uint32(data[len(data)-1]) << 8
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum = sum + (sum >> 16)
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
func createSocket()  int{
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		fmt.Printf("Socket creation error: %v\n", err)
		os.Exit(1)
	}
	return fd
}

func setPacket() []byte {
	packet := icmpPacket{
		Type:     8, // Echo request
		Code:     0,
		Checksum: 0,
		ID:       uint16(os.Getpid() & 0xffff),
		Sequence: 1,
		Data:     []byte("HELLO-ping"),
	}
	packetBytes := make([]byte, 8+len(packet.Data))
	packetBytes[0] = packet.Type
	packetBytes[1] = packet.Code
	binary.BigEndian.PutUint16(packetBytes[2:4], packet.Checksum)
	binary.BigEndian.PutUint16(packetBytes[4:6], packet.ID)
	binary.BigEndian.PutUint16(packetBytes[6:8], packet.Sequence)
	copy(packetBytes[8:], packet.Data)

	// Calculate and set checksum
	packet.Checksum = calculateChecksum(packetBytes)
	binary.BigEndian.PutUint16(packetBytes[2:4], packet.Checksum)
	return packetBytes
}

func sendPacket(ip string) {
	// Send packet
	fd := createSocket()
	defer syscall.Close(fd)
	packetBytes := setPacket()
	dst := &syscall.SockaddrInet4 {
		Port: 0,
	}
	dstIP := net.ParseIP(ip).To4() // Example: Google DNS
	copy(dst.Addr[:], dstIP)
	err := syscall.Sendto(fd, packetBytes, 0, dst)
	if err != nil {
		fmt.Printf("Send error: %v\n", err)
		os.Exit(1)
	}
	// Set read timeout
	tv := syscall.Timeval{
		Sec:  10, 
		Usec: 0,
	}
	err = syscall.SetsockoptTimeval(fd, syscall.SOL_SOCKET, syscall.SO_RCVTIMEO, &tv)
	if err != nil {
		fmt.Printf("Setting timeout error: %v\n", err)
		os.Exit(1)
	}
	receivePacket(fd)
}

func receivePacket(fd int) {
	// Receive reply
	reply := make([]byte, 1500)
	n, _, err := syscall.Recvfrom(fd, reply, 0)
	if err != nil {
		fmt.Printf("Receive error: %v\n", err)
		os.Exit(1)
	}
	// Parse reply (skip IP header)
	ipHeaderLen := int(reply[0]&0x0f) * 4
	icmpReply := reply[ipHeaderLen:n]
	RenderPacket(icmpReply)
}
