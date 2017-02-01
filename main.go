package main
import (
	"fmt"
	"os"
	"github.com/torukita/go-udpgen/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

/*
import (
	"github.com/torukita/go-udpgen/pkt"
	"github.com/torukita/go-udpgen/pcap"
	"log"
	"os"
	"strconv"
)

func main() {
	nic := os.Args[1]
	num, _ := strconv.Atoi(os.Args[2])
	handle, err := pcap.Open(nic)
	if err != nil {
			log.Fatal(err)
	}
	defer handle.Close()

	eth := pkt.NewEthernet("11:22:33:44:55:66", "88:77:66:55:44:33")
	ip := pkt.NewIPv4("10.10.0.10", "1.2.3.4")
	udp := pkt.NewUDP(8000, 5000)
	packet, err := pkt.UDPPacket(eth, ip, udp)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < num; i++ {
		pcap.Send(handle, packet)
	}
	log.Println("Done Send")
}
*/
