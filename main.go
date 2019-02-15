package main

import (
	"bufio"
	"fmt"
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

func main() {
	args := []string{"-A", "-f", "-o", "wq"}
	cmd := exec.Command("ps", args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		// Here map each line of output to a Process
		// Also get child processes for each process
		fmt.Println(scanner.Text())
		fmt.Println(strings.Fields(scanner.Text()))
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
