package api

import(
	"time"
	"runtime"
	"fmt"
	"net"
	"context"
	"github.com/torukita/go-udpgen/pcap"
	"github.com/torukita/go-udpgen/pkt"	
)

type Config struct {
	Device    string        `json:"device"`
	Core      int           `json:"core"`
	Count     uint64         `json:"count"`
	Second    time.Duration `json:"second"`
	SrcEth    string        `json:"srceth"`
	DstEth    string        `json:"dsteth"`
	SrcIP     string        `json:"srcip"`
	DstIP     string        `json:"srcdst"`
	SrcPort   uint16         `json:"srcport"`
	DstPort   uint16         `json:"dstport"`
}

func NewConfig() *Config {
	return &Config{
		Device: "eth0",
		Core: runtime.NumCPU(),
		SrcEth: "00:00:00:00:00:01",
		DstEth: "00:00:00:00:00:02",
		SrcIP: "10.0.0.1",
		DstIP: "10.0.0.2",
		SrcPort: 8888,
		DstPort: 5000,
	}
}

func (c *Config)String() string {
	str := fmt.Sprintf("Device: %s Core: %v SrcEth: %s DstEth: %s SrcIP: %s DstIP: %s SrcPort: %v DstPort: %v",
		c.Device, c.Core, c.SrcEth, c.DstEth, c.SrcIP, c.DstIP, c.SrcPort, c.DstPort)
	return str
}

func (c *Config)Dump() {
	fmt.Printf("%v\n", c)
}

func (c *Config)Exec() error {
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

func (c *Config)parse() error {
	if _, err := net.ParseMAC(c.SrcEth); err != nil {
		return err
	}
	if _, err := net.ParseMAC(c.DstEth); err != nil {
		return err
	}
	if net.ParseIP(c.SrcIP) == nil || net.ParseIP(c.DstIP) == nil {
		return fmt.Errorf("Could not parse IP")
	}
	return nil
}
