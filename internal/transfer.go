package internal

import (
	"errors"
	"os"
	"os/exec"
)

func Transfer(source Location, destination Location) error {
	if err := exec.Command("rsync", "--version").Run(); err != nil {
		return errors.New("rsync is not installed or not in the PATH")
	}

	if err := exec.Command("sshpass").Run(); err != nil {
		return errors.New("sshpass is not installed or not in the PATH")
	}

	if !source.IsLocal && !destination.IsLocal {
		return errors.New("todo: move between servers")
	}

	var password = ""
	if source.IsLocal {
		password = destination.Server.Password
	} else {
		password = source.Server.Password
	}

	rsync := exec.Command("sshpass", "-eRSYNC_PASSWORD", "rsync", "-ravzhP", source.RsyncString(), destination.RsyncString())
	rsync.Env = append(os.Environ(), "RSYNC_PASSWORD="+password)
	println(rsync.String())

	rsync.Stdout = os.Stdout
	rsync.Stderr = os.Stderr

	return rsync.Run()
}
