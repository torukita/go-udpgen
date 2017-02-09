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
	packet, err := pkt.UDPPacket(c.Size, eth, ip, udp)

	if err != nil {
		return err
	}

	d := worker.NewPacketSender(c.Concurrency, c.Device)

	d.Start()

	timer := time.NewTimer(time.Duration(c.Timer) * time.Second)
	if c.Timer == 0 { timer.Stop() }

	start := time.Now()	
Loop:
	for i := uint64(0); i < c.Count; i++ {
		select {
		case <-timer.C:
			break Loop
		default:
			d.Send(packet)
		}
	}
	timer.Stop()
	d.Stop()
	fmt.Printf("Required Time: %f sec\n", time.Since(start).Seconds())	
	return nil
}

