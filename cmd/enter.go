package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"gtan/internal"
	"os"
	"os/exec"
)

var EnterCmd = &cobra.Command{
	Use:   "enter",
	Short: "join to server via ssh",
	Long:  ".",
	Args:  cobra.MinimumNArgs(1),
	Run:   enter,
}

func init() {
}

func enter(cmd *cobra.Command, args []string) {
	serverName := args[0]
	server := func() internal.Server {
		for name, obj := range internal.UserData.Servers {
			if name == serverName {
				return obj
			}
		}

		cobra.CheckErr(errors.New("server not found"))
		return internal.Server{}
	}()

	if err := exec.Command("ssh").Run(); err != nil && errors.Is(err, exec.ErrNotFound) {
		cobra.CheckErr(errors.New("ssh is not installed or not in the PATH"))
	}

	if err := exec.Command("sshpass").Run(); err != nil {
		cobra.CheckErr(errors.New("sshpass is not installed or not in the PATH"))
	}

	ssh := exec.Command("sshpass", "-eRSYNC_PASSWORD", "ssh", server.Username+"@"+server.Hostname)
	ssh.Env = append(os.Environ(), "RSYNC_PASSWORD="+server.Password)
	println(ssh.String())

	ssh.Stdout = os.Stdout
	ssh.Stderr = os.Stderr
	ssh.Stdin = os.Stdin

	cobra.CheckErr(ssh.Run())
}
