package udp

import (
	"github.com/torukita/go-udpgen/pcap"
	"log"
)

var (
	device string = "en0"
)

func Send(packet []byte) error {
	handle, err := pcap.Open(device)
	if err != nil {
		return err
	}
	defer handle.Close()

	err = handle.WritePacketData(packet)
	if err != nil {
		log.Fatal("Could not send packet", err)
	}
	return err
}
