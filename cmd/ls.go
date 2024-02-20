package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gtan/internal"
	"os"
	"text/tabwriter"
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list server and named locations",
	Long:  ".",
	Run:   ls,
}

func init() {
	LsCmd.Flags().BoolP("server", "s", true, "server")
	LsCmd.Flags().BoolP("location", "l", true, "location")
}

func ls(cmd *cobra.Command, args []string) {
	isServer, _ := cmd.Flags().GetBool("server")
	isLocation, _ := cmd.Flags().GetBool("location")

	isServerSet := cmd.Flags().Changed("server")
	isLocationSet := cmd.Flags().Changed("location")

	if isServerSet || isLocationSet {
		isServer = false
		isLocation = false

		if isServerSet {
			isServer = true
		}

		if isLocationSet {
			isLocation = true
		}
	}

	if isServer {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

		fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s", "Name", "Host", "Username"))

		for serverName, server := range internal.UserData.Servers {
			fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s", serverName, server.Hostname, server.Username))
		}

		w.Flush()

		println()
	}

	if isLocation {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

		fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s", "Name", "Server", "Path"))

		for locationName, location := range internal.UserData.NamedLocations {
			fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s", locationName, location.ServerName, location.Path))
		}

		w.Flush()
	}
}
