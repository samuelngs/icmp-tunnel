package main

// interfaces is the type of interface
type interfaces string

// Interface types
var (
	InterfaceTypeTap  interfaces = "tap"
	InterfaceTypeTun  interfaces = "tun"
	InterfaceTypeEth  interfaces = "eth"
	InterfaceTypeVLan interfaces = "vlan"
	InterfaceTypeBr   interfaces = "br"
	InterfaceTypeWLan interfaces = "wlan"
	InterfaceTypeAth  interfaces = "ath"
)
