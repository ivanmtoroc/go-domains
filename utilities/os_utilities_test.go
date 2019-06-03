package utilities

import (
	"testing"
)

func TestRunShCommand(t *testing.T) {
	if stdout, _ := RunShCommant("echo 'Hire me!'"); stdout != "Hire me!" {
		t.Error("command execution failed")
	}
}
