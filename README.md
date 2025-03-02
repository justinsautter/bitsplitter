# BitSplitter

BitSplitter is a CLI-based, IPv4 network subnet calculator written in Golang.

## Installation

The BitSplitter executable file is included for Unix based systems (Linux/MacOS).

You can install it on your system and use BitSplitter from the terminal/command line anywhere doing the following:

`mv ./bitsplitter /usr/local/bin/bitsplitter`

I will be working on an auto installation for the future to cut out this step.

## Usage

Enter in the IP address with CIDR (Classless Inter-Domain Routing) subnet mask you wish to create a subnet for.

Ex: 172.16.27.0/18

```
bitsplitter 172.16.27.0/18
```

Output:
```
# Overview (IP/CIDR)
IP address: 172.16.0.0/18

# Address Details
Address range: 172.16.0.0 - 172.16.63.255
Number of hosts: 16384

Network address: 172.16.0.0
Broadcast address: 172.16.63.255

Usable range: 172.16.0.1 - 172.16.63.254
Usable hosts: 16382

# Mask information
IP subnet mask (binary): 11111111.11111111.11000000.00000000
Wildcard mask (binary): 00000000.00000000.00111111.11111111

# Classification
IP type: Private
IP class: B
```

![Usage Example Screenshot](https://github.com/user-attachments/assets/a4401865-a9a4-4f0c-8e61-838d3efebadb)

## Future plans

I plan to eventually make this compatible with IPv6 as well. I'm pretty new to this whole programming thing, so this might take awhile.

If there are any bugs to report, please let me know!