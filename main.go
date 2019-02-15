package main

import (
	p "github.com/leepuppychow/processes-and-threads/processes"
)

func main() {
	p.ExecuteCommand("ps", "-A", "-f", "-o", "wq")
}
