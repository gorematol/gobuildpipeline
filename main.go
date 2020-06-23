package main

import (
	"gobuildpipeline/builddocker"
	"gobuildpipeline/utils"
)

func main() {
	builddocker.Install()
        utils.PullImage("jenkins/jenkins")
}
