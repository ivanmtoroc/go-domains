package utilities

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

// RunShCommant execute bash command and return the stdout
func RunShCommant(command string) (string, error) {
	var output bytes.Buffer
	// -c flag mean that input is a command
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Stdout = &output

	if err := cmd.Run(); err != nil {
		log.Println("bash command execution failed")
		log.Println("- command: ", command)
		log.Println("- error: ", err)
		return "", err
	}
	return strings.TrimSuffix(output.String(), "\n"), nil
}
