package builddocker

import (
	"fmt"
	"sort"

	"gobuildpipeline/utils"
)

var installcmds = map[string]string{
	"Installing docker from yum channels": "yum install -y docker",
	"Checking binary using docker cmd":   "docker",
	"Starting docker from systemctl":   "systemctl start docker",
	"Checking running docker process":  "ps -ef | grep -i docker | grep -v grep",
}

func sortKeys(m map[string]string) []string {
	var orderedKeys []string
	for k, _ := range m {
		orderedKeys = append(orderedKeys, k)
	}
	sort.Strings(orderedKeys)
	fmt.Println(orderedKeys)
	return orderedKeys
}

func Install() {
	mapkeys := sortKeys(installcmds)
	for _, msg := range mapkeys {
		cmd := installcmds[msg]
		fmt.Println(msg)
		_, rc := utils.OSrunCmd(cmd)
		if rc != 0 {
		        fmt.Println("Failed OS command " + cmd)
			utils.Log.Fatal("OS command failed " + cmd, rc)
		}
	}
}
