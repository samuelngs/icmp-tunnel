package main

// InterfaceType is the type of interface
type InterfaceType struct {
	name string
}

// String returns interface type name
func (t *InterfaceType) String() string {
	return t.name
}

// Interface types
var (
	InterfaceTypeTap  = &InterfaceType{"tap"}
	InterfaceTypeTun  = &InterfaceType{"tun"}
	InterfaceTypeEth  = &InterfaceType{"eth"}
	InterfaceTypeVLan = &InterfaceType{"vlan"}
	InterfaceTypeBr   = &InterfaceType{"br"}
	InterfaceTypeWLan = &InterfaceType{"wlan"}
	InterfaceTypeAth  = &InterfaceType{"ath"}
)
