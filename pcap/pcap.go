package pcap

import (
	"github.com/google/gopacket/pcap"
	"time"
	"log"
)

var (
	device       string        = "en0"
	snapshot_len int32         = 1024
	promiscuous  bool          = false
	timeout      time.Duration = 30 * time.Second
)

func Open(device string) (*pcap.Handle, error) {
			handle, err := pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
			if err != nil {
				return nil, err
			}
			return handle, nil
}

func Close(handle *pcap.Handle) {
	handle.Close()
}

func Send(handle *pcap.Handle, packet []byte) error {
	err := handle.WritePacketData(packet)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
