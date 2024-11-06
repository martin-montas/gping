package ICMPRequest

import (
        "fmt"
        "net"
        "os"
)

func ReadPing(conn net.Conn) {
	// Read the response
	buf := make([]byte, 1500)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		os.Exit(1)
	}

	fmt.Printf("Received %d bytes: %v\n", n, buf[:n])
}

func SendPing(conn net.Conn) {
        // Send an ICMP echo request
		var err error
        _, err = conn.Write([]byte{8, 0, 0, 0, 0, 0, 0, 0})
        if err != nil {
                fmt.Println("Error writing to connection:", err)
                os.Exit(1)
        }
}

func Run(ip *string) {
	// Create a new ICMP connection
	conn, err := net.Dial("ip4:icmp", *ip)
	if err != nil {
		fmt.Println("Error dialing:", err)
		os.Exit(1)
	}
	defer conn.Close()

	SendPing(conn)
	ReadPing(conn)
}

