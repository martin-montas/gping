package ICMPRequest

import (
    "encoding/binary"
    "syscall"
    "strings"
    "fmt"
    // "net"
    "os"
    // "syscall"
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
		// sendPacket(ip)
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


func sendPacket() {
    packet := icmpPacket {
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

    // Resolve destination address
    dst := &syscall.SockaddrInet4{
        Port: 0,
    }
}
