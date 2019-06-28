package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	indent := strings.Repeat(" ", len("Interface "))
	for _, i := range interfaces {
		fmt.Printf("Interface Name: %v\n", i.Name)
		fmt.Printf("%sFlags: %v\n", indent, i.Flags.String())
		fmt.Printf("%sMTU: %v\n", indent, i.MTU)
		fmt.Printf("%sHardware Address: %v\n", indent, i.HardwareAddr)
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
		}
		addresses, err := byName.Addrs()
		for k, v := range addresses {
			fmt.Printf("%sAddress #%v: %v\n", indent, k, v.String())
		}
		fmt.Println()
	}
}
