package utils

import (
	"bytes"
	"os/exec"
)

func OSrunCmd(oscmd string) (string, string, int) {
	var shell = "bash"
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(shell, "-c", oscmd)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// https://stackoverflow.com/a/55055100
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			errStr := string(stderr.Bytes())
			return "", errStr, exitError.ExitCode()
		}
	}
	outStr := string(stdout.Bytes())
	return outStr, nil, 0
}
