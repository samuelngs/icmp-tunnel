package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func addRoute(dest, route, iface string) IError {

	scmd := fmt.Sprintf("ip -4 r a %s via %s dev %s", dest, route, iface)
	cmd := exec.Command("sh", "-c", scmd)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func delRoute(dest string) IError {

	sargs := fmt.Sprintf("-4 route del %s", dest)
	args := strings.Split(sargs, " ")
	cmd := exec.Command("ip", args...)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
