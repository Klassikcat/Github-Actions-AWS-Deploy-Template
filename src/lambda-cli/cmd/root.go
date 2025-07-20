package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "poorserverless",
	Short: "PoorServerless - A simple serverless function generator",
	Long: `PoorServerless is a CLI tool that helps you create serverless functions
with templates and configuration files for AWS Lambda deployment.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(createFunctionCmd)
}
