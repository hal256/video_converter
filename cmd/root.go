package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	cfgFile string
)

func init() {
	rootCmd.PersistentFlags()
}

var rootCmd = &cobra.Command{
	Use:   "videoc",
	Short: "Video Converter",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
