package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/torukita/go-udpgen/api"
	"log"
	"strconv"
	"time"
	_ "context"
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
		config := api.NewConfig()
		config.Device = viper.GetString("root.device")
		config.SrcEth = viper.GetString("root.src-eth")
		config.DstEth = viper.GetString("root.dst-eth")
		config.SrcIP = viper.GetString("root.src-ip")
		config.DstIP = viper.GetString("root.dst-ip")		
		udp_src, _ := strconv.Atoi(viper.GetString("root.src-port"))
		udp_dst, _ := strconv.Atoi(viper.GetString("root.dst-port"))
		config.SrcPort = uint16(udp_src)
		config.DstPort = uint16(udp_dst)
		config.Second = time.Duration(viper.GetInt64("root.time")) * time.Second
		config.Count = uint64(viper.GetInt64("root.count"))

		err := config.Exec()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Done Send")
	},
}

func init() {
	RootCmd.AddCommand(arootCmd)
	RootCmd.AddCommand(versionCmd)

	arootCmd.Flags().String("src-eth", "00:00:00:00:00:01", "Source mac address")
	arootCmd.Flags().String("dst-eth", "00:00:00:00:00:02", "Dest mac address")
	arootCmd.Flags().String("src-ip", "10.0.40.1", "Source IP address")
	arootCmd.Flags().String("dst-ip", "10.0.40.2", "Dest IP address")
	arootCmd.Flags().String("src-port", "9999", "UDP source port")
	arootCmd.Flags().String("dst-port", "9999", "UDP dest port")
	arootCmd.Flags().String("device", "eth0", "Device name")
	arootCmd.Flags().Uint64("time", 0, "seconds which keeps sending packtes")
	arootCmd.Flags().Uint64("count", 1, "The number of packets to be send")
	

	viper.BindPFlag("root.src-eth", arootCmd.Flags().Lookup("src-eth"))
	viper.BindPFlag("root.dst-eth", arootCmd.Flags().Lookup("dst-eth"))
	viper.BindPFlag("root.src-ip", arootCmd.Flags().Lookup("src-ip"))
	viper.BindPFlag("root.dst-ip", arootCmd.Flags().Lookup("dst-ip"))
	viper.BindPFlag("root.src-port", arootCmd.Flags().Lookup("src-port"))
	viper.BindPFlag("root.dst-port", arootCmd.Flags().Lookup("dst-port"))
	viper.BindPFlag("root.device", arootCmd.Flags().Lookup("device"))
	viper.BindPFlag("root.time", arootCmd.Flags().Lookup("time"))
	viper.BindPFlag("root.count", arootCmd.Flags().Lookup("count"))

}
