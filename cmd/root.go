package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tardigo",
	Short: "an ETL tool designed to get data from anywhere to anywhere",
	Long: `
	tardigo is a developer friendly, plugin based ETL tool that supports extensibility.
	refer https://github.com/tardigodev/tardigo.
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
