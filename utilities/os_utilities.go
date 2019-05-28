package utilities

import (
  "log"
  "bytes"
  "os/exec"
  "strings"
)

// Function to execute bash command
func RunShCommant(command string) string {
  var out bytes.Buffer

  // -c flag mean that
  cmd := exec.Command("/bin/bash", "-c", command)
  cmd.Stdout = &out

  if err := cmd.Run(); err != nil {
    log.Println("bash command execution failed")
    log.Println("- command: ", command)
    log.Fatalln("- error: ", err)
  }
  return strings.TrimSuffix(out.String(), "\n")
}
