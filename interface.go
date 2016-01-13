package main

import (
	"os"
	"strings"
	"syscall"
	"unsafe"
)

// IfReq represents interface request
type IfReq struct {
	Name  [0x10]byte
	Flags uint16
	pad   [0x28 - 0x10 - 2]byte
}

// Interface - the network interfaces
type Interface struct {
	name   string
	class  *InterfaceType
	device *os.File
}

// Name returns Interface name
func (i *Interface) Name() string {
	return i.name
}

// Class returns Interface type
func (i *Interface) Class() *InterfaceType {
	return i.class
}

// NewInterface creates new interface
func NewInterface(name string, class *InterfaceType, device string, flags uint16) (*Interface, IError) {
	file, err := os.OpenFile(device, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}
	ifr := &IfReq{}
	ifr.Flags = flags
	copy(ifr.Name[:], name)
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, file.Fd(), uintptr(syscall.TUNSETIFF), uintptr(unsafe.Pointer(&ifr))); err != 0 {
		return nil, err
	}
	ifName := strings.Trim(string(ifr.Name[:]), "\x00")
	networkif := &Interface{class: class, name: ifName, device: file}
	return networkif, nil
}

// TapInterface creates tap interface
func TapInterface() (*Interface, IError) {
	return NewInterface("tap", InterfaceTypeTap, "/dev/net/tun", 0x0002|0x1000)
}

// TunInterface creates tun interface
func TunInterface() (*Interface, IError) {
	return NewInterface("tun", InterfaceTypeTun, "/dev/net/tun", 0x0001|0x1000)
}
