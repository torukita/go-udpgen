package util

import (
	"log"
	"net"
)

type Device struct {
	Name    string
	MacAddr string
}

type Devices []Device

func CollectDevice() Devices {
	devices := make(Devices, 0)

	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range ifaces {
		devices = append(devices, Device{
			Name:    i.Name,
			MacAddr: fmt.Sprintf("%s", i.HardwareAddr),
		})
	}
	return devices
}
