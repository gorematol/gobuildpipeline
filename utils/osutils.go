package utils

import (
	"bytes"
	"os/exec"
 	"fmt"
)

var shell = "bash"

func OSrunCmd(oscmd string) (string, error) {
 	fmt.Println("Running OS command:-", oscmd)
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(shell, "-c", oscmd)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	outStr := string(stdout.Bytes())
	if len(outStr) == 0 {
		return outStr, nil
	}
	if err != nil {
                return "", err     
	}
	return outStr, nil
}

