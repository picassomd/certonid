package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	PlatformName = ""
	Version      = "unknown-version"
	GitCommit    = "unknown-commit"
	BuildTime    = "unknown-buildtime"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Certonid",
	Long:  `All software has versions. This is Certonid's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Certonid v0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
