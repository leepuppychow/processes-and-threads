package processes

import (
	h "github.com/leepuppychow/processes-and-threads/helpers"
)

/*
	1. Get all processes (PIDs only)
	2. Iterate through and Get individual process info
	3. Create a Process struct
	4. Now get children of this process (pgrep -P [PID])
	5. For each child repeat steps 2-4
	6. If there are no more children then stop
*/

type Process struct {
	ProcessId   string
	ThreadCount int
	Children    map[string]*Process
	Command     string
	Duration    string
}

func MakeProcessTree() *Process {
	root := Process{
		ProcessId:   "root",
		ThreadCount: 0,
		Children:    make(map[string]*Process),
		Command:     "",
		Duration:    "",
	}
	allPIDs := h.ExecuteCommand("ps", "-A", "-o", "pid")
	for i, PID := range allPIDs {
		if i == 0 {
			continue
		}
		childProcess := CreateSingleProcess(PID)
		root.Children[PID] = childProcess
		if childProcess != nil {
			go childProcess.FindChildren()
		}
	}
	return &root
}

func CreateSingleProcess(PID string) *Process {
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
		childProcess := CreateSingleProcess(PID)
		p.Children[PID] = childProcess
		if childProcess != nil {
			childProcess.FindChildren()
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
			go c.ChildrenCount(result)
		}
	}
}
