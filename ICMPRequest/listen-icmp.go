package ICMPRequest

import (
	"fmt"
	"encoding/binary"
)

func RenderPacket(icmpReply []byte,ip string) {
	if icmpReply[0] == 0 { // ICMP Echo Reply
		fmt.Printf("IP: %s\n", ip)
		fmt.Printf("Got echo reply!\n")
		fmt.Printf("Type: %d\n", icmpReply[0])
		fmt.Printf("Code: %d\n", icmpReply[1])
		fmt.Printf("ID: %d\n", binary.BigEndian.Uint16(icmpReply[4:6]))
		fmt.Printf("Sequence: %d\n", binary.BigEndian.Uint16(icmpReply[6:8]))
		fmt.Printf("Data: %s\n", string(icmpReply[8:]))
	} else {
		fmt.Printf("IP: %s\n", ip)
		fmt.Printf("This is an Error: \n")
		fmt.Printf("Type: %d\n", icmpReply[0])
		fmt.Printf("Code: %d\n", icmpReply[1])
	}
}

