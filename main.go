package main

import (
	p "github.com/leepuppychow/processes-and-threads/processes"
)

func main() {
	// all := p.CreateProcessList(p.ExecuteCommand("ps", "-A", "-f", "-o", "wq"))
	// fmt.Println(all)
	// fmt.Println(len(all))
	root := p.InitializeRoot()
	p.PopulateProcessTree(root)
}
