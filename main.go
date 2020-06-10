package main

import (
	"fmt"

	"minikubeinstall/prechecks"
)

var cpucmd = "/bin/grep -E --color 'svm|vmx' /proc/cpuinfo"

func main() {
	cpu := prechecks.OSrunCmd(cpucmd)
	fmt.Println(cpu)
}
