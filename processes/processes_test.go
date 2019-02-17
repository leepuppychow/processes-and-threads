package processes

import "testing"

func TestCreateProcess(t *testing.T) {
	root := CreateProcess("1")
	if root.ProcessId != "1" {
		t.Errorf("Expected id of 1, got: %s", root.ProcessId)
	}
	if root.ThreadCount < 0 {
		t.Error("Expected 0 or more threads on root process' queue")
	}
	if root.Command == "" {
		t.Error("Root process should have some associated command")
	}
	if root.Duration == "" {
		t.Error("Root process should have some duration value")
	}
}
