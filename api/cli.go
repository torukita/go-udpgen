package api

import(
	"time"
	"fmt"
	"context"
	"github.com/torukita/go-udpgen/pcap"
	"github.com/torukita/go-udpgen/pkt"	
)

func (c *Config)ExecFromCLI() error {
	if err := c.parse(); err != nil {
		return err
	}
	handle, err := pcap.Open(c.Device)
	if err != nil {
		return err
	}
	defer func() {
		// fmt.Println("closing.... ")
		handle.Close()
	}()

	eth := pkt.NewEthernet(c.SrcEth, c.DstEth)
	ip  := pkt.NewIPv4(c.SrcIP, c.DstIP)
	udp := pkt.NewUDP(c.SrcPort, c.DstPort)
	packet, err := pkt.UDPPacket(eth, ip, udp)
	if err != nil {
		return err
	}

	start := time.Now()

	if (c.Second == 0) {
		for i := uint64(0); i < c.Count; i++ {
			if err := pcap.Send(handle, packet); err != nil {
				return err
			}
		}
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), c.Second)
		defer cancel()
		errCh := make(chan error, 1)
		counter := 0
		go func(ctx context.Context) {
			for {
				err := pcap.Send(handle, packet)
				if err != nil {
					errCh <- err
				}
				counter++
			}
		}(ctx)
		select {
		case <- ctx.Done():
			// fmt.Println("done:", ctx.Err())
			fmt.Println("Sent Packets: ", counter)
		case err := <- errCh:
			fmt.Println("Happend error in go routine:", err)
		}
	}
	end := time.Now()
	
	fmt.Printf("Required Time: %f sec\n",(end.Sub(start)).Seconds()) 	// Stop watch
	return nil
}

