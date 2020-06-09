package main

import (
        "fmt" 

	"minikubeinstall/prechecks"
)
var grepcmd = "/bin/grep"
var vmxcpuargs = []string{"-E", "--color", "vmx", "/proc/cpuinfo" }
var svmcpuargs = []string{"-E", "--color", "svm", "/proc/cpuinfo" }


func main() {
    vmx := prechecks.CheckCPU(grepcmd, vmxcpuargs)
    svm := prechecks.CheckCPU(grepcmd, svmcpuargs)

    fmt.Println(svm)
    fmt.Println(vmx)
}  
