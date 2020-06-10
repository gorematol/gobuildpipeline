package osutils

import (
	"bytes"
	"log"
	"os/exec"
)

var shell = "bash"

func OSrunCmd(oscmd string) string {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(shell, "-c", oscmd)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	if len(outStr) == 0 {
		return outStr
	}
	if err != nil {
		log.Fatal("cmd.Run() failed with %s\n", errStr)
	}
	return outStr
}

func DownloadFiles(url string) string {
	//download manifest files for processing
}
