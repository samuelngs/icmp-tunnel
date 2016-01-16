package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

// RunServer runs server-side daemon
func RunServer(c *cli.Context) {

	// creates a tap interface
	iface, err := TapInterface()

	if err != nil {
		log.Printf("creating interface error: %v\n", err)
	}

	iface.ParseIP(10, 0, 40, 1)
	iface.Up()

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

	fmt.Println(iface.Name())

	for {
		select {
		case buffer := <-data:
			if buffer != nil {
			}
			// fmt.Println(buffer[:])
		}
	}
}
