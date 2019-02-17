package processes

import (
	h "github.com/leepuppychow/processes-and-threads/helpers"
)

type Process struct {
	ProcessId   string
	ThreadCount int
	Children    map[string]*Process
	Command     string
	Duration    string
}

func MakeProcessTree() *Process {
	root := CreateProcess("1")
	root.FindChildren()
	return root
}

func CreateProcess(PID string) *Process {
	processInfo := h.ExecuteCommand("ps", "-o pid", "-o wq", "-o comm", "-o time", "-p", PID)[4:]
	if len(processInfo) == 0 {
		return nil
	}
	return &Process{
		ProcessId:   processInfo[0],
		ThreadCount: h.ToInt(processInfo[1]),
		Children:    make(map[string]*Process),
		Command:     processInfo[2],
		Duration:    processInfo[3],
	}
}

func (p *Process) FindChildren() {
	children := h.ExecuteCommand("pgrep", "-P", p.ProcessId)
	if len(children) == 0 {
		return
	}
	for _, PID := range children {
		childProcess := CreateProcess(PID)
		p.Children[PID] = childProcess
		if childProcess != nil {
			go childProcess.FindChildren()
		}
	}
}

func (p *Process) ChildrenCount(result map[string]int) {
	if len(p.Children) == 0 {
		return
	}
	result[p.ProcessId] = len(p.Children)
	for _, c := range p.Children {
		if c != nil {
			c.ChildrenCount(result)
		}
	}
}
