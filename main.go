package main

import (
	"bitsplitter/subnet"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: bitsplitter <IP/CIDR>")
		fmt.Println("Example: bitsplitter 192.168.0.0/24")
		os.Exit(1)
	}

	info, err := subnet.CalculateSubnet(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Print results
	fmt.Println("\n# Overview (IP/CIDR)")
	fmt.Printf("IP address: %s/%d\n", info.IP, info.CIDR)
	fmt.Println("\n# Address Details")
	fmt.Printf("Address range: %s\n", info.AddressRange)
	fmt.Printf("Number of hosts: %d\n\n", info.NumHosts)
	fmt.Printf("Network address: %s\n", info.NetworkAddr)
	fmt.Printf("Broadcast address: %s\n\n", info.BroadcastAddr)
	fmt.Printf("Usable range: %s\n", info.UsableRange)
	fmt.Printf("Usable hosts: %d\n", info.UsableHosts)
	fmt.Println("\n# Mask information")
	fmt.Printf("IP subnet mask (binary): %s\n", info.SubnetMaskBin)
	fmt.Printf("Wildcard mask (binary): %s\n", info.WildcardMaskBin)
	fmt.Println("\n# Classification")
	fmt.Printf("IP type: %s\n", map[bool]string{true: "Private", false: "Public"}[info.IsPrivate])
	fmt.Printf("IP class: %s\n", info.IPClass)
}
