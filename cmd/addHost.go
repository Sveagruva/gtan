package cmd

import (
	"github.com/spf13/cobra"
	"gtan/internal"
	"os"
)

var AddHostCmd = &cobra.Command{
	Use:   "addHost",
	Short: "Add a new host",
	Long:  ".",
	Args:  cobra.MinimumNArgs(1),
	Run:   addHost,
}

func init() {
	AddHostCmd.Flags().StringP("username", "u", "", "Username")
	AddHostCmd.Flags().StringP("hostname", "s", "", "Hostname")
	AddHostCmd.Flags().StringP("password", "p", "", "Password")
}

func addHost(cmd *cobra.Command, args []string) {
	name := args[0]

	hostname, _ := cmd.Flags().GetString("hostname")
	username, _ := cmd.Flags().GetString("username")
	password, _ := cmd.Flags().GetString("password")

	if hostname == "" {
		cmd.Help()
		os.Exit(1)
	}

	cobra.CheckErr(
		internal.UserData.AddHost(name, hostname, username, password),
	)
}
