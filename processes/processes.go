package processes

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	h "github.com/leepuppychow/processes-and-threads/helpers"
)

type Process struct {
	ProcessId   int
	ThreadCount int
	Children    map[string]*Process
	Command     string
	Duration    string
}

func ExecuteCommand(command string, args ...string) {
	output, err := exec.Command(command, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
	fmt.Println(CreateProcessList(string(output)))
}

func CreateProcessList(stdout string) []Process {
	allProcs := []Process{}
	for i, line := range strings.Split(strings.TrimSpace(stdout), "\n") {
		if i == 0 {
			continue
		}
		lineSlice := strings.Fields(line)
		p := Process{
			ProcessId:   h.ToInt(lineSlice[1]),
			ThreadCount: h.ToInt(lineSlice[8]),
			Children:    nil,
			Command:     lineSlice[7],
			Duration:    lineSlice[6],
		}
		allProcs = append(allProcs, p)
	}

	return allProcs
}
