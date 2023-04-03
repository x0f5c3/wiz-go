package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	"github.com/x0f5c3/wiz-go/net"
)

// discoverCmd represents the discover command
var discoverCmd = &cobra.Command{
	Use:  "discover",
	RunE: runE,
}

func runE(_ *cobra.Command, _ []string) error {
	proto, err := net.NewProtocol()
	if err != nil {
		return err
	}
	disco, err := proto.Discover()
	if err != nil {
		return err
	}
	pterm.Success.Printfln("Got msg %s", disco)
	return nil
}

func init() {
	rootCmd.AddCommand(discoverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// discoverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// discoverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
