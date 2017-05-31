package cmd

import (
	"github.com/spf13/cobra"
)

var LoongCmd = &cobra.Command{
	Use: "loong",
	Short: "loong manages your files",
	Long: `loong is the main command, used to manage your files.

Loong is a Simple, Fast and Strong file manager built in Go.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var LogLevel string

func init() {
	LoongCmd.Flags().StringVar(&LogLevel, "logLevel", "warn", "Log Level: debug, info, warn, error, fatal, panic.")
}
