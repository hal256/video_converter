package cmd

import (
	"github.com/hal256/video_converter/internal/hls"
	"github.com/spf13/cobra"
)

type Options struct {
	targetPath string
	distPath string
}
var (
	o = &Options{}
)

func init() {
	convertCmd.Flags().StringVarP(&o.targetPath, "target", "t", "./targetPath/", "for target")
	convertCmd.Flags().StringVarP(&o.distPath, "dist", "d", "./distPath/", "for dist ")
	rootCmd.AddCommand(convertCmd)
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert all file",
	Run: func(cmd *cobra.Command, args []string) {
		hls.ConvertAllFIle(o.targetPath, o.distPath)
	},
}
