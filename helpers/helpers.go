package helpers

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func ExecuteCommand(command string, args ...string) []string {
	output, err := exec.Command(command, args...).Output()
	if err != nil {
		log.Println(command, err)
	}
	return strings.Fields(string(output))
}

func ToInt(val string) int {
	num, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println(err)
		num = 0
	}
	return num
}
