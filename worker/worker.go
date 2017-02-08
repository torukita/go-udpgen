package worker

import(
	"sync"
	"github.com/torukita/go-udpgen/device"
	"fmt"
	"os"
)

var (
	maxQueues = 1000
)

type PacketSender struct {
	maxWorkers   int
	deviceHandle *device.Handle
	queue        chan []byte // packet
	wg           sync.WaitGroup
}

func NewPacketSender(maxWorkers int, devicename string) *PacketSender {
	handle, err := device.Open(devicename)
	if err != nil {
		fmt.Println("error")
		os.Exit(-1)
	}
	d := &PacketSender{
		maxWorkers: maxWorkers,
		deviceHandle: handle,
		queue: make(chan []byte, maxQueues),
	}
	return d
}

func (d *PacketSender) Send(packet []byte) {
	d.queue <- packet
}

func (d *PacketSender) Start() {
	d.wg.Add(d.maxWorkers)
	for i :=0; i < d.maxWorkers; i++ {
		go func() {
			defer d.wg.Done()
			for packet := range d.queue {
				if err := device.Send(d.deviceHandle, packet); err != nil {
					fmt.Println(err)
					return
				}
			}
		}()
	}
}

func (d *PacketSender) Stop() {
	defer func() {
		fmt.Println("closing...")
		device.Close(d.deviceHandle)
	}()
	close(d.queue)
	d.wg.Wait()
}
