package api

import(
	"time"
	"fmt"
	"github.com/torukita/go-udpgen/pkt"
	"github.com/torukita/go-udpgen/worker"
)

func (c *Config)ExecFromCLI() error {
	if err := c.parse(); err != nil {
		return err
	}

	eth := pkt.NewEthernet(c.SrcEth, c.DstEth)
	ip  := pkt.NewIPv4(c.SrcIP, c.DstIP)
	udp := pkt.NewUDP(c.SrcPort, c.DstPort)
	packet, err := pkt.UDPPacket(eth, ip, udp)

	if err != nil {
		return err
	}

	d := worker.NewPacketSender(c.Concurrency, c.Device)

	d.Start()
	start := time.Now()
	for i := uint64(0); i < c.Count; i++ {
		d.Send(packet)
	}
	end := time.Now()
	fmt.Printf("Required Time: %f sec\n",(end.Sub(start)).Seconds())
	d.Stop()
	return nil
}

