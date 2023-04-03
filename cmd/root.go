package cmd

import (
	"os"
	"os/signal"

	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "wiz-go",
	Version: "v0.0.5", // <---VERSION---> Updating this version, will also create a new GitHub release.
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Fetch user interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		pterm.Warning.Println("user interrupt")
		checkForUpdates()
		os.Exit(0)
	}()

	// Execute cobra
	if err := rootCmd.Execute(); err != nil {
		checkForUpdates()
		os.Exit(1)
	}

	checkForUpdates()
}

func checkForUpdates() {
	err := pcli.CheckForUpdates()
	pterm.Fatal.PrintOnError(err)
}

func setRepo(repo string) {
	err := pcli.SetRepo(repo)
	pterm.Fatal.PrintOnError(err)
}

func init() {
	// Adds global flags for PTerm settings.
	// Fill the empty strings with the shorthand variant (if you like to have one).
	rootCmd.PersistentFlags().BoolVarP(&pterm.PrintDebugMessages, "debug", "d", false, "enable debug messages")
	rootCmd.PersistentFlags().BoolVar(&pterm.RawOutput, "raw", false, "print unstyled raw output (set it if output is written to a file)")
	rootCmd.PersistentFlags().BoolVar(&pcli.DisableUpdateChecking, "no-update", false, "disables update checks")

	// Use https://github.com/pterm/pcli to style the output of cobra.
	setRepo("x0f5c3/wiz-go")
	pcli.SetRootCmd(rootCmd)
	pcli.Setup()

	// Change global PTerm theme
	pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgCyan)
}
