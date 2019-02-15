package processes

import (
	"log"
	"os/exec"
	"strings"

	h "github.com/leepuppychow/processes-and-threads/helpers"
)

type Process struct {
	ProcessId   int
	ThreadCount int
	Children    map[int]*Process
	Command     string
	Duration    string
}

/*

	1. Get all processes (id only)
	2. Iterate through and Get individual process info (ps -f -p [PID] -o wq)
	3. Create a Process struct
	4. Now get children of this process (pgrep -P [PID])
	5. For each child repeat steps 2-4
	6. If there are no more children then stop

*/

func ExecuteCommand(command string, args ...string) []string {
	output, err := exec.Command(command, args...).Output()
	if err != nil {
		log.Println(command, err)
	}
	return strings.Fields(string(output))
}

func InitializeRoot() *Process {
	root := Process{
		ProcessId:   0,
		ThreadCount: 0,
		Children:    make(map[int]*Process),
		Command:     "",
		Duration:    "",
	}
	allPIDs := ExecuteCommand("ps", "-A", "-o", "pid")
	for i, PID := range allPIDs {
		if i == 0 {
			continue
		}
		root.Children[h.ToInt(PID)] = CreateSingleProcess(PID)
	}
	return &root
}

func CreateSingleProcess(PID string) *Process {
	processInfo := ExecuteCommand("ps", "-f", "-o", "wq", "-p", PID)[9:]
	if len(processInfo) == 0 {
		return nil
	}
	return &Process{
		ProcessId:   h.ToInt(processInfo[1]),
		ThreadCount: h.ToInt(processInfo[8]),
		Children:    nil,
		Command:     processInfo[7],
		Duration:    processInfo[6],
	}
}
