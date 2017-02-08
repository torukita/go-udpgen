package api

import(
	"time"
	"fmt"
	"net/http"
	"context"
	"github.com/labstack/echo"	
	"github.com/torukita/go-udpgen/pkt"
	"github.com/torukita/go-udpgen/worker"
)

var (
	procSend int = 0
)

func WebSend(c echo.Context) error {
	req := NewConfig()
	if err := c.Bind(req); err != nil {
		return err
	}
	fmt.Printf("%+v", req)

	if procSend > 0 {
		return c.JSON(http.StatusBadRequest, nil)
	}

	err := req.ExecFromWeb()
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusAccepted, nil)
}

func (c *Config)ExecFromWeb() error {
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
	for i := uint64(0); i < c.Count; i++ {
		d.Send(packet)
	}
	d.Stop()
	return nil
}

func Send(handle interface{}, packet []byte) error {
//	fmt.Println(time.Now())
	return nil
}

func SendTimer(timeout time.Duration, handle interface{}, packet []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer func() {
		fmt.Println("called defer cancel")
		cancel()
	}()
	errCh := make(chan error, 1)
	stop := 0
	go func () {
		for  {
			if stop == 1 { break }
			if err := Send(handle, packet); err != nil {
				errCh <- err
				break
			}
		}
		return
	}()
	select {
	case <-ctx.Done():
		fmt.Println("called done:", ctx.Err())
		stop = 1
	case err := <-errCh:
		fmt.Println(err)
	}
	return nil
}


