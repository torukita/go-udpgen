package worker

import(
	"sync"
	"github.com/torukita/go-udpgen/device"
	"fmt"
	"os"
)

type PacketSender struct {
	maxWorkers   int
	maxQueues    int64
	deviceHandle *device.Handle
	queue        chan []byte // packet
	wg           sync.WaitGroup
}

func NewPacketSender(maxWorkers int, maxQueues int64, devicename string) *PacketSender {
	handle, err := device.Open(devicename)
	if err != nil {
		fmt.Println("error")
		os.Exit(-1)
	}
	d := &PacketSender{
		maxWorkers: maxWorkers,
		maxQueues: maxQueues,
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
	close(d.queue)
	d.wg.Wait()
	defer func() {
		fmt.Println("closing...")
		device.Close(d.deviceHandle)
	}()
}
