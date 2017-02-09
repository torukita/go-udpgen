package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = "v0.0.6"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Long:  `go-udpgen application version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

var RootCmd = &cobra.Command{
	Use:   "go-udpgen",
	Short: "udp generator for golang",
	Long:  `go-udpgen can be used to send UDP packets from CLI or WEB`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
