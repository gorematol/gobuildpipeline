package main

import (
	"fmt"

	"minikubeinstall/utils"
)

var cpucmd = "/bin/grep -E --color 'svm|vmx' /proc/cpuinfo"

func main() {
	cpu := osutils.OSrunCmd(cpucmd)
	fmt.Println(cpu)
}
