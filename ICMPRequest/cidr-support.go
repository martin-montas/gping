package ICMPRequest

import (
	"strings"
	"fmt"
	"net"
	"strconv"
)

func isCidrVAlidIpv6(ip string) bool {
	// /0 to /128 are valid ranges.
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
	// /0 to /32 are valid ranges.
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
		sendSingleIP(ip.String())
	}
}

func handleCidr4(ip string) { 
	_, ipNet, err := net.ParseCIDR(ip)
	if err != nil {
		fmt.Println("Error parsing CIDR:", err)
		return
	}
	firstIP := ipNet.IP
	for ip := firstIP.Mask(ipNet.Mask); ipNet.Contains(ip); incrementIP(ip) {
		sendSingleIP(ip.String())
	}
}
