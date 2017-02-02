package pcap

import (
	"github.com/google/gopacket/pcap"
	"time"
)

var (
	snapshot_len int32         = 64
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
	if err := handle.WritePacketData(packet); err != nil {
		return err
	}
	return nil
}

func SendOne(handle *pcap.Handle, packet []byte) error {
	if err := handle.WritePacketData(packet); err != nil {
		return err
	}
	return nil
}
