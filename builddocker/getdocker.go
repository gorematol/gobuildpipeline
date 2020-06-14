package builddocker

import (
	"fmt"

	"gobuildpipeline/utils"
)

var orderedkeys = []string

var installcmds = map[string]string{
	"Installing docker": "yum install -y docker",
	"Checking binary":   "docker",
	"Starting docker":   "systemctl docker start",
	"Checking process":  "ps -ef | grep -i docker",
}

func sortKeys(key string) () {
}

func Install() {
	for msg, cmd := range installcmds {
		fmt.Println(msg + ":- " + cmd)
		_, err := utils.OSrunCmd(cmd)
		if err != nil {
			utils.Log.Fatal("cmd.Run() failed with %s\n", err)
		}
	}
}
