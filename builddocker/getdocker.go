package builddocker

import (
	"fmt"
	"sort"

	"gobuildpipeline/utils"
)

var dockerinstall = map[string]string{
	"1. Installing docker from yum channels": "yum install -y docker",
	"2. Checking binary using docker cmd": "docker",
	"3. Enabling docker": "systemctl enable docker",
	"4. Starting docker": "systemctl start docker",
	"5. Checking running docker process": "ps -ef | grep -i docker | grep -v grep",
}

func sortKeys(m map[string]string) []string {
	var orderedKeys []string
	for k, _ := range m {
		orderedKeys = append(orderedKeys, k)
	}
	sort.Strings(orderedKeys)
	return orderedKeys
}

func Install() {
	mapkeys := sortKeys(dockerinstall)
	for _, msg := range mapkeys {
		cmd := dockerinstall[msg]
		fmt.Println(msg)
		out, err, rc := utils.OSrunCmd(cmd)
		if rc != 0 {
			fmt.Println("Failed:- " + cmd)
			utils.Log.Fatal("OS command failed "+cmd, rc, err)
		} else {
			utils.Log.Print(out)
		}
	}
}

