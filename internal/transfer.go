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

	if !source.IsLocal && !destination.IsLocal {
		return errors.New("todo: move between servers")
	}

	// todo: password
	rsync := exec.Command("rsync", "-ravzP", source.RsyncString(), destination.RsyncString())
	println(rsync.String())

	rsync.Stdout = os.Stdout
	rsync.Stderr = os.Stderr

	return rsync.Run()
}
