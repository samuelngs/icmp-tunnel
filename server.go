package main

import (
	"fmt"

	"golang.org/x/net/icmp"

	"github.com/codegangsta/cli"
)

// RunServer runs server-side daemon
func RunServer(c *cli.Context) {

	iface, err := CreateIface()
	if err != nil {
		panic(fmt.Sprintf("creating interface error: %v\n", err))
	}
	defer iface.Close()

	conn, err := ListenICMP()
	if err != nil {
		panic(fmt.Sprintf("listen icmp packet error: %v\n", err))
	}
	defer conn.Close()

	data := make(chan []byte, 8)

	go func() {
		for {
			buffer := make([]byte, 1522)
			_, err := iface.Read(buffer)
			if err == nil {
				data <- buffer
			}
		}
	}()

	for {
		select {
		case buffer := <-data:
			fmt.Println(string(buffer[:]))
		}
	}
}

// CreateIface creates network interface
func CreateIface() (*Interface, error) {

	// creates a tap interface
	iface, err := TapInterface()

	if err == nil {
		// set network interface ip
		iface.ParseIP(10, 0, 40, 1)
		iface.Up()
	}

	return iface, err
}

// ListenICMP listens to incoming ICMP packet
func ListenICMP() (*icmp.PacketConn, error) {

	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	defer func() {
		conn.Close()
	}()

	return conn, err
}
