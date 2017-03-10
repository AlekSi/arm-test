package main

import (
	"flag"
	"fmt"
	"runtime"
)

func main() {
	flag.Parse()
	fmt.Printf("Hello from %s %s!\n", runtime.GOOS, runtime.GOARCH)
}
