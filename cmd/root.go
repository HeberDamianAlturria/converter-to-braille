package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "converter-to-braille",
	Short: "A simple braille converter",
	Long:  "A simple command line tool to convert images to braille",
	Run:   runRoot,
}

func runRoot(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
