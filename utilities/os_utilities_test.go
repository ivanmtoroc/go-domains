package utilities

import (
  "testing"
)

func TestRunShCommand(t *testing.T) {
  if stdout := RunShCommant("echo 'Hire me!'"); stdout != "Hire me!" {
    t.Error("Command execution failed")
  }
}
