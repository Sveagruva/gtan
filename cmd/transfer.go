package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	. "gtan/internal"
	"os"
)

var TransferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "transfer files",
	Long:  ".",
	Args:  cobra.MinimumNArgs(1),
	Run:   transfer,
}

func init() {
}

func transfer(cmd *cobra.Command, args []string) {
	source := args[0]
	destination := "_default_destination"
	if len(args) > 1 {
		destination = args[1]
	}

	if source == "" {
		cmd.Help()
		os.Exit(1)
	}

	sourceLocation, err := ParseLocation(source)
	cobra.CheckErr(err)
	destinationLocation, err := ParseLocation(destination)
	cobra.CheckErr(err)

	sourceLocation.Path, err = HydratePath(sourceLocation.Path)
	cobra.CheckErr(err)
	destinationLocation.Path, err = HydratePath(destinationLocation.Path)
	cobra.CheckErr(err)

	if sourceLocation.IsLocal && destinationLocation.IsLocal {
		cobra.CheckErr(errors.New("use cp or mv for local operations"))
	}

	cobra.CheckErr(
		Transfer(sourceLocation, destinationLocation),
	)
}
