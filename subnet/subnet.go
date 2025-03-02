package subnet

import (
	"fmt"
	"net"
)

// SubnetInfo holds all calculated subnet info
type SubnetInfo struct {
	IP              net.IP
	CIDR            int
	AddressRange    string
	NumHosts        uint32
	UsableRange     string
	NetworkAddr     net.IP
	BroadcastAddr   net.IP
	UsableHosts     uint32
	SubnetMaskBin   string
	WildcardMaskBin string
	IPClass         string
	IsPrivate       bool
}

// Computes subnet details from a CIDR notation
func CalculateSubnet(input string) (*SubnetInfo, error) {
	// Parse the IP and CIDR
	_, ipNet, err := net.ParseCIDR(input)
	if err != nil {
		return nil, fmt.Errorf("invalid CIDR notation: %v", err)
	}

	// Get IP and mask
	ip := ipNet.IP
	mask := ipNet.Mask
	cidr, _ := mask.Size()

	// Calculate number of hosts
	totalBits := 32
	hostBits := totalBits - cidr
	numHosts := uint32(1 << hostBits)

	// Calculate network and broadcast addresses
	network := ip.Mask(mask)
	broadcast := make(net.IP, len(ip))
	for i := range ip {
		broadcast[i] = network[i] | ^mask[i]
	}

	// Calculate usable hosts
	var usableHosts uint32
	var usableRange string
	if numHosts >= 2 {
		usableHosts = numHosts - 2
		usableFirst := incrementIP(network)
		usableLast := decrementIP(broadcast)
		usableRange = fmt.Sprintf("%s - %s", usableFirst, usableLast)
	} else {
		usableRange = "N/A (too small for usable hosts)"
	}

	// Convert masks to binary
	subnetMaskBin := maskToBinary(mask)
	wildcardMaskBin := maskToBinary(invertMask(mask))

	// Determine IP class
	ipClass := getIPClass(ip)

	// Create SubnetInfo
	info := &SubnetInfo{
		IP:              ip,
		CIDR:            cidr,
		AddressRange:    fmt.Sprintf("%s - %s", network, broadcast),
		NumHosts:        numHosts,
		UsableRange:     usableRange,
		NetworkAddr:     network,
		BroadcastAddr:   broadcast,
		UsableHosts:     usableHosts,
		SubnetMaskBin:   subnetMaskBin,
		WildcardMaskBin: wildcardMaskBin,
		IPClass:         ipClass,
		IsPrivate:       ip.IsPrivate(),
	}

	return info, nil
}

// Converts mask to binary string format
func maskToBinary(mask net.IPMask) string {
	binary := ""
	for i, b := range mask {
		binary += fmt.Sprintf("%08b", b)
		if i < len(mask)-1 {
			binary += "."
		}
	}
	return binary
}

// Creates a wildcard mask by inverting the subnet mask
func invertMask(mask net.IPMask) net.IPMask {
	inverted := make(net.IPMask, len(mask))
	for i, b := range mask {
		inverted[i] = ^b
	}
	return inverted
}

// Increments an IP address by 1
func incrementIP(ip net.IP) net.IP {
	newIP := make(net.IP, len(ip))
	copy(newIP, ip)
	for i := len(newIP) - 1; i >= 0; i-- {
		if newIP[i] < 255 {
			newIP[i]++
			break
		}
		newIP[i] = 0
	}
	return newIP
}

// Decrements an IP address by 1
func decrementIP(ip net.IP) net.IP {
	newIP := make(net.IP, len(ip))
	copy(newIP, ip)
	for i := len(newIP) - 1; i >= 0; i-- {
		if newIP[i] > 0 {
			newIP[i]--
			break
		}
		newIP[i] = 255
	}
	return newIP
}

// Determines the IP class (A, B, C, D, E) based on the first octet
func getIPClass(ip net.IP) string {
	firstOctet := ip[0]
	switch {
	case firstOctet <= 127:
		return "A"
	case firstOctet <= 191:
		return "B"
	case firstOctet <= 223:
		return "C"
	case firstOctet <= 239:
		return "D"
	default:
		return "E"
	}
}
