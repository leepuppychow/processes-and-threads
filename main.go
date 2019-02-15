package main

import (
	"fmt"
	"time"

	p "github.com/leepuppychow/processes-and-threads/processes"
)

func main() {
	start := time.Now()
	root := p.MakeProcessTree()
	fmt.Println(root.Children["300"])
	fmt.Println(time.Since(start))
}
