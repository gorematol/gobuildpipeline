package prechecks

import (
	"os/exec"
	"log"
)


func CheckCPU(oscmd string, args []string) (string)  {
        cmd := exec.Command(oscmd, args...)
        out, err := cmd.CombinedOutput()
        var empty string

        if err != nil {
		log.Fatal("bad apple")
	}
        if string(out) == empty {
 		log.Fatal("Virtualization is not supported on this host")
	}
        
        return string(out) 
}
        

