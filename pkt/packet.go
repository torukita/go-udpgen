package pkt

import (
	"github.com/google/gopacket"
	"bytes"
	"log"
)

func UDPPacket(framesize int, eth *Ethernet, ip *IPv4, udp *UDP) ([]byte, error) {
	options := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	// eth header length: 14
	// ip header length: 20
	// udp header length: 8
	// udp payload lenght: 18
	// fcs : 4
	// Total: 64 bytes (small packet)
	//	rawBytes := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}

	//rawBytes := make([]byte, 18)
	payloadLength := framesize - 46
	rawBytes := bytes.Repeat([]byte{1}, payloadLength)

	if err := udp.udp.SetNetworkLayerForChecksum(ip.ipv4); err != nil {
		return nil, err
	}
	buffer := gopacket.NewSerializeBuffer()
	err := gopacket.SerializeLayers(buffer, options,
		eth.eth,
		ip.ipv4,
		udp.udp,
		gopacket.Payload(rawBytes),
	)
	if err != nil {
		log.Fatal("hoge", err)
		return nil, err
	}
	out := buffer.Bytes()
	return out, nil
}
