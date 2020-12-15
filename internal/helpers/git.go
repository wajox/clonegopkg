package helpers

import (
	"os"
	"os/exec"
)

func CloneGitRepository(repository, directory string) error {
	cmd := exec.Command("git", "clone", repository, directory)

	return cmd.Run()
}

func RemoveGitRemoteOrigin(directory string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	defer func() {
		_ = os.Chdir(wd)
	}()

	if err := os.Chdir(directory); err != nil {
		return err
	}

	cmd := exec.Command("git", "remote", "remove", "origin")

	return cmd.Run()
}
