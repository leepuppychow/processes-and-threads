package main

import (
	"fmt"
	"time"

	p "github.com/leepuppychow/processes-and-threads/processes"
)

func main() {
	start := time.Now()
	root := p.MakeProcessTree()
	fmt.Println(time.Since(start))

	procsWithChildrenCount := make(map[string]int)
	root.ChildrenCount(procsWithChildrenCount)
	fmt.Println(procsWithChildrenCount)

	//test
}
