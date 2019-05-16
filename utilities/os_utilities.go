package utilities

import (
  "log"
  "bytes"
  "os/exec"
  "strings"
)

func RunShCommant(command string) string {
	cmd := exec.Command("/bin/sh", "-c", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Println("Command execution failed:")
		log.Println(command)
		log.Fatalln(err)
	}
	return strings.TrimSuffix(out.String(), "\n")
}
