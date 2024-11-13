package ICMPRequest

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)


func isCidrVAlidIpv6(ip string) bool {
	lastTwoChars := ip[len(ip)-2:]
	if !strings.ContainsRune(ip,'/') {
		return false
	}
	if strings.Index(string(lastTwoChars[0]), "/") != 0 {
		num, err := strconv.Atoi(lastTwoChars)
		if err != nil {
			return false
		}
		if num > 128 {
			return false
		}
		return true
	}
	return false
}

func isCidrVAlidIpv4(ip string) bool {
	lastTwoChars := ip[len(ip)-2:]
	if !strings.ContainsRune(ip,'/') {
		return false
	}
	if strings.Index(string(lastTwoChars[0]), "/") != 0 {
		num, err := strconv.Atoi(lastTwoChars)
		if err != nil {
			return false
		}
		if num > 32 {
			return false
		}
		return true
	}
	return true 
}

func handleCidr6(ip string) { 
	_, ipNet, err := net.ParseCIDR(ip)
	if err != nil {
		fmt.Println("Error parsing CIDR:", err)
		return
	}
	firstIP := ipNet.IP
	for ip := firstIP.Mask(ipNet.Mask); ipNet.Contains(ip); incrementIP(ip) {
		sendPacket(ip.String())
	}
}

func handleCidr4(ip string) { 
	_, ipNet, err := net.ParseCIDR(ip)
	fmt.Printf("ipNet: %v\n", ipNet)
	if err != nil {
		fmt.Println("Error parsing CIDR:", err)
		return
	}
	firstIP := ipNet.IP
	for ip := firstIP.Mask(ipNet.Mask); ipNet.Contains(ip); incrementIP(ip) {
		sendPacket(ip.String())
	}
}

func incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
