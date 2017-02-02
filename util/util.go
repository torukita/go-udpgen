package util

import (
	"net"
	"fmt"
	"github.com/vishvananda/netlink"
)

var (
	DeviceList func() Devices
)

type Device struct {
	Index   int
	Name    string
	MacAddr string
	Type    string
}

type Devices []Device

type Arp struct {
	IP      string
	MacAddr string
}

type ArpTable []Arp

func linkListFromNetPackage() Devices {
	devices := make(Devices, 0)
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		if i.Name == "lo" {
			continue
		}
		devices = append(devices, Device{
			Index:   i.Index,
			Name:    i.Name,
			MacAddr: fmt.Sprintf("%s", i.HardwareAddr),
		})
	}
	return devices
}

func linkList() Devices {
	devices := make(Devices, 0)
	links, _ := netlink.LinkList()
	for _, v := range(links) {
		if v.Attrs().Name == "lo" {
			continue
		}
		devices = append(devices, Device{
			Index: v.Attrs().Index,
			Name:  v.Attrs().Name,
			MacAddr: fmt.Sprintf("%s", v.Attrs().HardwareAddr),
			Type: v.Type(),
		})
	}
	return devices
}

func ArpList(linkIndex int) ArpTable {
	table := make(ArpTable, 0)
	// NeighList(LinkIndex, Family)
	t, _ := netlink.NeighList(linkIndex, netlink.FAMILY_V4)
	for _, v := range(t) {
		table = append(table, Arp{
			IP: fmt.Sprintf("%s", v.IP),
			MacAddr: fmt.Sprintf("%s", v.HardwareAddr),
		})
	}
	return table
}

func init() {
	DeviceList = linkList
	//DeviceList = linkListFromNetPackage
}
