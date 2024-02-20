package cmd

import (
	"github.com/spf13/cobra"
	"gtan/internal"
)

var RemoveLocationCmd = &cobra.Command{
	Use:   "removeLocation",
	Short: "remove location",
	Long:  ".",
	Args:  cobra.MinimumNArgs(1),
	Run:   removeLocation,
}

func init() {
}

func removeLocation(cmd *cobra.Command, args []string) {
	name := args[0]

	cobra.CheckErr(
		internal.UserData.RemoveNamedLocation(name),
	)
}
