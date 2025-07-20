package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lambda-cli",
	Short: "Lambda CLI - A modern serverless function generator",
	Long: `Lambda CLI is a tool that helps you create serverless functions
with spec.yaml configuration and AWS CDK deployment infrastructure.

Based on GEMINI.md specification for simplified Lambda deployment.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(createFunctionCmd)
}
