package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "1.0.0"
var rootCmd = &cobra.Command{
	Use:     "snorlax",
	Version: version,
	Short:   "snorlax - Start snorlax webserver",
	Long: `snorlax - Start snorlax webserver
   
It provides some functionalities like:
* When is the next train goes to city center

Maintainer: Son La
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'\n", err)
		os.Exit(1)
	}
}
