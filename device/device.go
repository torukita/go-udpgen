package device

import (
	"github.com/google/gopacket/pcap"
	"time"
)

var (
	snapshot_len int32         = 64
	promiscuous  bool          = false
	timeout      time.Duration = 30 * time.Second
)

type Handle struct {
	handle *pcap.Handle
}

func Open(device string) (*Handle, error) {
	handle, err := pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		return nil, err
	}
	h := &Handle{
		handle: handle,
	}
	return h, nil
}

func Close(h *Handle) {
	h.handle.Close()
}

func Send(h *Handle, packet []byte) error {
	if err := h.handle.WritePacketData(packet); err != nil {
		return err
	}
	return nil
}

func SendOne(h *Handle, packet []byte) error {
	if err := h.handle.WritePacketData(packet); err != nil {
		return err
	}
	return nil
}
