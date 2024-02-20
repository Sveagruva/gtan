package cmd

import (
	"github.com/spf13/cobra"
	"gtan/internal"
)

var AddLocationCmd = &cobra.Command{
	Use:   "addLocation",
	Short: "Add a new location",
	Long:  ".",
	Args:  cobra.MinimumNArgs(3),
	Run:   addLocation,
}

func init() {
}

func addLocation(cmd *cobra.Command, args []string) {
	name := args[0]
	server := args[1]
	path := args[2]

	cobra.CheckErr(
		internal.UserData.AddNamedLocation(name, server, path),
	)
}
