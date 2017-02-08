package api

import(
	"time"
	"runtime"
	"fmt"
	"net"
)

var (
	allowedNet = []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"}
)

func IsAllowedNet(s string) bool {
	for _, v := range allowedNet {
		_, ipnet, _ := net.ParseCIDR(v)
		if ipnet.Contains(net.ParseIP(s)) {
			return true
		}
	}
	return false
}

type Config struct {
	Device      string         `json:"device"`
	Concurrency int            `json:"concurrency"`
	Core        int            `json:"core"`
	Count       uint64         `json:"count"`
	Second      time.Duration  `json:"second"`
	SrcEth      string         `json:"srceth"`
	DstEth      string         `json:"dsteth"`
	SrcIP       string         `json:"srcip"`
	DstIP       string         `json:"dstip"`
	SrcPort     uint16         `json:"srcport"`
	DstPort     uint16         `json:"dstport"`
	Size        int            `json:"size"`
}

func NewConfig() *Config {
	// initial param is changed in the later
	return &Config{
		Device: "eth0",
		Core: runtime.NumCPU(),
		Concurrency: runtime.GOMAXPROCS(runtime.NumCPU()),
		SrcEth: "00:00:00:00:00:01",
		DstEth: "00:00:00:00:00:02",
		SrcIP: "10.0.0.1",
		DstIP: "10.0.0.2",
		SrcPort: 8888,
		DstPort: 5000,
		Size: 64,
	}
}

func (c *Config)Dump() {
	fmt.Printf("%+v\n", c)
}

func (c *Config)parse() error {
	if _, err := net.ParseMAC(c.SrcEth); err != nil {
		return err
	}
	if _, err := net.ParseMAC(c.DstEth); err != nil {
		return err
	}
	if !IsAllowedNet(c.SrcIP) || !IsAllowedNet(c.DstIP) {
		return fmt.Errorf("Denied IP")
	}
	if c.Second != 0 {
		c.Second = c.Second * time.Second
	}
	return nil
}

