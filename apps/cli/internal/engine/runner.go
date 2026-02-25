package engine

import (
	"os"
	"os/exec"
)

func ExecutePill(path string, args []string) error {
	cmd := exec.Command(path, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	cmd.Env = os.Environ()

	return cmd.Run()
}
