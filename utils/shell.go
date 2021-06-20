package utils

import "os/exec"

func Command(command string) string {
	cmd := exec.Command("/bin/bash", "-c", command)
	stdout, _ := cmd.Output()
	return string(stdout)
}
