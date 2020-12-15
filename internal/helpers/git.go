package helpers

import (
	"os/exec"
)

func CloneGitRepository(repository, directory string) error {
	cmd := exec.Command("git", "clone", repository, directory)

	return cmd.Run()
}
