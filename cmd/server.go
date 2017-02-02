package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/torukita/go-udpgen/server"
)

var serverCmd = &cobra.Command {
	Use: "server",
	Short: "Run API&Web server",
	Long: `API & Web Server`,
	Run: func(cmd *cobra.Command, args []string) {
		addr := viper.GetString("server.listen") + ":" + viper.GetString("server.port")
		server.Run(addr, viper.GetBool("server.debug"))
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)

	serverCmd.Flags().String("listen", "0.0.0.0", "The listen ip of the server")
	serverCmd.Flags().String("port", "9000", "The port of the server")
	serverCmd.Flags().Bool("debug", false, "debug mode for the server")

	viper.BindPFlag("server.port", serverCmd.Flags().Lookup("port"))
	viper.BindPFlag("server.listen", serverCmd.Flags().Lookup("listen"))
	viper.BindPFlag("server.debug", serverCmd.Flags().Lookup("debug"))
}
