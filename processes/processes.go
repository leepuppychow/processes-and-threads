package processes

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

type Process struct {
	ProcessId   int
	ThreadCount int
	Children    []*Process
	Command     string
	Duration    string
}

func ExecuteCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	go PrintOutput(stdout)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

func PrintOutput(stdout io.ReadCloser) {
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		// Here map each line of output to a Process
		// Also get child processes for each process
		fmt.Println(scanner.Text())
		fmt.Println(strings.Fields(scanner.Text()))
	}
}
