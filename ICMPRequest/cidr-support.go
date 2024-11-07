package ICMPRequest

import (
	"strings"
	"fmt"
	"net"
	"strconv"
)

func isCidrVAlidIpv4(ip string) bool {
	// /0 to /32 are valid ranges.
	lastTwo := ip[len(ip)-2:]
	if !strings.ContainsRune(ip,'/') {
		return false
	}
	if strings.Index(string(lastTwo[0]), "/") != 0 {
		num, err := strconv.Atoi(lastTwo)
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

func handleCidr(ip string) { 
	_, ipNet, err := net.ParseCIDR(ip)
	if err != nil {
		fmt.Println("Error parsing CIDR:", err)
		return
	}
	firstIP := ipNet.IP
	for ip := firstIP.Mask(ipNet.Mask); ipNet.Contains(ip); incrementIP(ip) {
		sendIP(ip.String())
	}
}
