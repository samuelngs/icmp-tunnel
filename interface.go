package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
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
	class  interfaces
	device *os.File
	ipnet  net.IPNet
}

// Name returns Interface name
func (i *Interface) Name() string {
	return i.name
}

// Class returns Interface type
func (i *Interface) Class() interfaces {
	return i.class
}

// Write implements io.Writer interface.
func (i *Interface) Write(p []byte) (int, error) {
	return i.device.Write(p)
}

// Read implements io.Reader interface.
func (i *Interface) Read(p []byte) (int, error) {
	return i.device.Read(p)
}

// ParseIP parses ip address
func (i *Interface) ParseIP(a, b, c, d byte) {
	self := net.IPv4(a, b, c, d)
	mask := self.DefaultMask()
	i.ipnet = net.IPNet{IP: self, Mask: mask}
}

// Up brings interface up/online
func (i *Interface) Up() error {
	if err := exec.Command("ip", "link", "set", i.name, "up").Run(); err != nil {
		return err
	}
	if err := exec.Command("ip", "addr", "add", i.ipnet.String(), "dev", i.name).Run(); err != nil {
		return err
	}
	return nil
}

// NewInterface creates new interface
func NewInterface(name string, class interfaces, device string, flags uint16) (*Interface, IError) {
	file, err := os.OpenFile(device, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}
	ifr := IfReq{}
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
func TapInterface(names ...string) (*Interface, IError) {
	var name string
	for _, val := range names {
		name = val
		break
	}
	return NewInterface(name, InterfaceTypeTap, "/dev/net/tun", 0x0002|0x1000)
}

// TunInterface creates tun interface
func TunInterface(names ...string) (*Interface, IError) {
	var name string
	for _, val := range names {
		name = val
		break
	}
	return NewInterface(name, InterfaceTypeTun, "/dev/net/tun", 0x0001|0x1000)
}

// IPAddRoute to add route via ip command
func IPAddRoute(dest, route, iface string) IError {
	scmd := fmt.Sprintf("ip -4 r a %s via %s dev %s", dest, route, iface)
	cmd := exec.Command("sh", "-c", scmd)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// IPDeleteRoute to remove route via ip command
func IPDeleteRoute(dest string) IError {
	sargs := fmt.Sprintf("-4 route del %s", dest)
	args := strings.Split(sargs, " ")
	cmd := exec.Command("ip", args...)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
