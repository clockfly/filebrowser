package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// Execute executes the commands.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func RootCmd() *cobra.Command {
	return rootCmd
}
