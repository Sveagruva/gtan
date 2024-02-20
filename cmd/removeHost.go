package cmd

import (
	"github.com/spf13/cobra"
	"gtan/internal"
)

var RemoveHostCmd = &cobra.Command{
	Use:   "removeHost",
	Short: "Remove host",
	Long:  ".",
	Args:  cobra.MinimumNArgs(1),
	Run:   removeHost,
}

func init() {
}

func removeHost(cmd *cobra.Command, args []string) {
	name := args[0]

	cobra.CheckErr(
		internal.UserData.RemoveHost(name),
	)
}
