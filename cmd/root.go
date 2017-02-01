package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/torukita/go-udpgen/pcap"
	"github.com/torukita/go-udpgen/pkt"
	"log"
	"net"
	"strconv"
	"time"
	"context"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Long:  `All software has versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.1")
	},
}

var RootCmd = &cobra.Command{
	Use:   "go-udpgen",
	Short: "go-udpgen golang",
	Long:  `This program is example collections for golagn`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("command root")
	},
}


var arootCmd = &cobra.Command{
	Use:   "hoge",
	Short: "udp generator for golang",
	Long:  `UDP packets generate from cmd or web`,
	Run: func(cmd *cobra.Command, args []string) {
		nic := viper.GetString("root.device")
		handle, err := pcap.Open(nic)
		if err != nil {
			log.Fatal(err)
		}
		defer handle.Close()
    timeout, _ := strconv.Atoi(viper.GetString("root.time"))

		if _, err := net.ParseMAC(viper.GetString("root.eth-src")); err != nil {
			log.Fatal(err)
		}
		if _, err := net.ParseMAC(viper.GetString("root.eth-dst")); err != nil {
			log.Fatal(err)
		}
		eth := pkt.NewEthernet(viper.GetString("root.eth-src"), viper.GetString("root.eth-dst"))

		if net.ParseIP(viper.GetString("root.ip-src")) == nil {
			log.Fatal()
		}
		if net.ParseIP(viper.GetString("root.ip-dst")) == nil {
			log.Fatal()
		}
		ip := pkt.NewIPv4(viper.GetString("root.ip-src"), viper.GetString("root.ip-dst"))

		udp_src, _ := strconv.Atoi(viper.GetString("root.udp-src"))
		udp_dst, _ := strconv.Atoi(viper.GetString("root.udp-dst"))
		udp := pkt.NewUDP(uint16(udp_src), uint16(udp_dst))
		packet, err := pkt.UDPPacket(eth, ip, udp)
		if err != nil {
			log.Fatal(err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)* time.Second)
		defer cancel()
		go func(ctx context.Context) {
			for {
				pcap.Send(handle, packet)
			}
		}(ctx)

		select {
		case <- ctx.Done():
			log.Println("done:", ctx.Err())
		}
		log.Println("Done Send")
	},
}

func init() {
	RootCmd.AddCommand(arootCmd)
	RootCmd.AddCommand(versionCmd)

	arootCmd.Flags().String("eth-src", "00:00:00:00:00:01", "Source mac address")
	arootCmd.Flags().String("eth-dst", "00:00:00:00:00:02", "Dest mac address")
	arootCmd.Flags().String("ip-src", "10.0.40.1", "Source IP address")
	arootCmd.Flags().String("ip-dst", "10.0.40.2", "Dest IP address")
	arootCmd.Flags().String("udp-src", "9999", "UDP source port")
	arootCmd.Flags().String("udp-dst", "9999", "UDP dest port")
	arootCmd.Flags().String("device", "eth0", "Device name")
  arootCmd.Flags().String("time", "2", "Timeout")

	viper.BindPFlag("root.eth-src", arootCmd.Flags().Lookup("eth-src"))
	viper.BindPFlag("root.eth-dst", arootCmd.Flags().Lookup("eth-dst"))
	viper.BindPFlag("root.ip-src", arootCmd.Flags().Lookup("ip-src"))
	viper.BindPFlag("root.ip-dst", arootCmd.Flags().Lookup("ip-dst"))
	viper.BindPFlag("root.udp-src", arootCmd.Flags().Lookup("udp-src"))
	viper.BindPFlag("root.udp-dst", arootCmd.Flags().Lookup("udp-dst"))
	viper.BindPFlag("root.device", arootCmd.Flags().Lookup("device"))
	viper.BindPFlag("root.time", arootCmd.Flags().Lookup("time"))

}
