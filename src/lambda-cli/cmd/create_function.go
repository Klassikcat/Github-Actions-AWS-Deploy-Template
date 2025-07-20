package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"poorserverless/internal/generator"

	"github.com/spf13/cobra"
)

var (
	functionName string
	runtime      string
	outputDir    string
)

var createFunctionCmd = &cobra.Command{
	Use:   "create-function",
	Short: "Create a new serverless function",
	Long: `Create a new serverless function with the specified runtime and configuration.
Supported runtimes: python, nodejs, go`,
	RunE: createFunction,
}

func init() {
	createFunctionCmd.Flags().StringVarP(&functionName, "name", "n", "", "Function name (required)")
	createFunctionCmd.Flags().StringVarP(&runtime, "runtime", "r", "python", "Runtime (python, nodejs, go)")
	createFunctionCmd.Flags().StringVarP(&outputDir, "output", "o", ".", "Output directory")
	createFunctionCmd.MarkFlagRequired("name")
}

func createFunction(cmd *cobra.Command, args []string) error {
	if functionName == "" {
		return fmt.Errorf("function name is required")
	}

	// Validate runtime
	supportedRuntimes := []string{"python", "nodejs", "go"}
	if !contains(supportedRuntimes, runtime) {
		return fmt.Errorf("unsupported runtime: %s. Supported runtimes: %s", runtime, strings.Join(supportedRuntimes, ", "))
	}

	// Create function directory
	functionDir := filepath.Join(outputDir, functionName)
	if err := os.MkdirAll(functionDir, 0755); err != nil {
		return fmt.Errorf("failed to create function directory: %w", err)
	}

	// Generate function files
	generator := generator.New(runtime, functionName, functionDir)
	if err := generator.Generate(); err != nil {
		return fmt.Errorf("failed to generate function: %w", err)
	}

	fmt.Printf("✅ Successfully created function '%s' with runtime '%s' in '%s'\n", functionName, runtime, functionDir)
	fmt.Printf("📁 Function directory: %s\n", functionDir)

	return nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
