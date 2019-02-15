package processes

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
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
	CreateProcessList(string(output))
}

func CreateProcessList(stdout string) []Process {
	allProcs := []Process{}
	stringLines := strings.Split(strings.TrimSpace(stdout), "\n")
	for i, line := range stringLines {
		if i == 0 {
			continue
		}
		lineSlice := strings.Fields(line)
		id, err := strconv.Atoi(lineSlice[1])
		if err != nil {
			fmt.Println(err)
		}

		threadCount, err := strconv.Atoi(lineSlice[8])
		if err != nil {
			fmt.Println(err)
			threadCount = 0
		}

		p := Process{
			ProcessId:   id,
			ThreadCount: threadCount,
			Children:    nil,
			Command:     lineSlice[7],
			Duration:    lineSlice[6],
		}
		allProcs = append(allProcs, p)
	}

	return allProcs
}
