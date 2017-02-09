package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/torukita/go-udpgen/api"
	"log"
	"strconv"
	"time"
	"os"
)

var sendCmd = &cobra.Command{
	Use:   "send [Interface Name]",
	Short: "Send UDP packets",
	Long:  `go-udpgen send can be used to send UDP packates from CLI`,
	Example: `$ go-udpgen send eth0 --dst-ip 10.10.10.10`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			os.Exit(-1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		devicename := args[0]
		config := api.NewConfig()
		config.Device = devicename
		config.SrcEth = viper.GetString("send.src-eth")
		config.DstEth = viper.GetString("send.dst-eth")
		config.SrcIP = viper.GetString("send.src-ip")
		config.DstIP = viper.GetString("send.dst-ip")		
		udp_src, _ := strconv.Atoi(viper.GetString("send.src-port"))
		udp_dst, _ := strconv.Atoi(viper.GetString("send.dst-port"))
		config.SrcPort = uint16(udp_src)
		config.DstPort = uint16(udp_dst)
		config.Second = time.Duration(viper.GetInt64("send.time")) * time.Second
		config.Count = uint64(viper.GetInt64("send.count"))
		if viper.GetInt("send.concurrency") != 0 {
			config.Concurrency = viper.GetInt("send.concurrency")
		}
		if viper.GetInt("send.size") != 64 {
			config.Size = viper.GetInt("send.size")
		}

		err := config.ExecFromCLI()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Done Send")
	},
}

func init() {
	RootCmd.AddCommand(sendCmd)

	sendCmd.Flags().String("src-eth", "00:00:00:00:00:01", "Source mac address")
	viper.BindPFlag("send.src-eth", sendCmd.Flags().Lookup("src-eth"))

	sendCmd.Flags().String("dst-eth", "00:00:00:00:00:02", "Dest mac address")
	viper.BindPFlag("send.dst-eth", sendCmd.Flags().Lookup("dst-eth"))

	sendCmd.Flags().String("src-ip", "10.0.40.1", "Source IP address")
	viper.BindPFlag("send.src-ip", sendCmd.Flags().Lookup("src-ip"))

	sendCmd.Flags().String("dst-ip", "10.0.40.2", "Dest IP address")
	viper.BindPFlag("send.dst-ip", sendCmd.Flags().Lookup("dst-ip"))

	sendCmd.Flags().String("src-port", "9999", "UDP source port")
	viper.BindPFlag("send.src-port", sendCmd.Flags().Lookup("src-port"))

	sendCmd.Flags().String("dst-port", "9999", "UDP dest port")
	viper.BindPFlag("send.dst-port", sendCmd.Flags().Lookup("dst-port"))

	sendCmd.Flags().Uint64("time", 0, "seconds which keeps sending packtes")
	viper.BindPFlag("send.time", sendCmd.Flags().Lookup("time"))

	sendCmd.Flags().Uint64("count", 1, "The number of packets to be send")
	viper.BindPFlag("send.count", sendCmd.Flags().Lookup("count"))	

	sendCmd.Flags().Int("concurrency", 0, "The number of goroutines to use")
	viper.BindPFlag("send.concurrency", sendCmd.Flags().Lookup("concurrency"))

	sendCmd.Flags().Int("size", 64, "The length of ethernet frame (size >=64)")
	viper.BindPFlag("send.size", sendCmd.Flags().Lookup("size"))

}
