package cmd

import (
	"gtan/internal"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gtan",
	Short: "GTan is a SSH transfer utility on top of Rsync",
	Long:  ``,
	Run:   nil,
}

func Execute() {
	cobra.CheckErr(internal.UserData.Load())

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(TransferCmd)
	rootCmd.AddCommand(EnterCmd)
	rootCmd.AddCommand(LsCmd)
	rootCmd.AddCommand(AddHostCmd)
	rootCmd.AddCommand(RemoveHostCmd)
	rootCmd.AddCommand(AddLocationCmd)
	rootCmd.AddCommand(RemoveLocationCmd)
}
