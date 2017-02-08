package api

import(
	"time"
	"fmt"
	"net/http"
	"context"
	"github.com/labstack/echo"	
	"github.com/torukita/go-udpgen/device"
	"github.com/torukita/go-udpgen/pkt"	
)

var (
	procSend int = 0
)

func WebSend(c echo.Context) error {
	req := new(Config)
	if err := c.Bind(req); err != nil {
		return err
	}

	if procSend > 0 {
		return c.JSON(http.StatusBadRequest, nil)
	}
	procSend++
//	fmt.Printf("Start(%v)", procSend)	
	err := req.ExecFromWeb()
	procSend--
//	fmt.Printf("Stop(%v)", procSend)
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	//return c.JSON(http.StatusAccepted, nil)
	return c.JSON(http.StatusOK, nil)	
}

func (c *Config)ExecFromWeb() error {
	if err := c.parse(); err != nil {
		return err
	}
	handle, err := device.Open(c.Device)
	if err != nil {
		return err
	}
	defer func() {
		fmt.Println("handle closed")
		device.Close(handle)
	}()

	eth := pkt.NewEthernet(c.SrcEth, c.DstEth)
	ip  := pkt.NewIPv4(c.SrcIP, c.DstIP)
	udp := pkt.NewUDP(c.SrcPort, c.DstPort)
	packet, err := pkt.UDPPacket(eth, ip, udp)
	if err != nil {
		return err
	}

	if (c.Second == 0) {
		go func() {
			for i:= uint64(0); i < c.Count; i++ {
				if err := Send(handle,packet); err != nil {
					fmt.Println(err)
					return
				}
			}
			fmt.Println("Done send packet")
		}()
	} else {
		go SendTimer(c.Second, handle, packet)
	}
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


