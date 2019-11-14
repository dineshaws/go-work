package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Print OS and ARCH")
	fmt.Println("OS ", runtime.GOOS)
	fmt.Println("ARCH ", runtime.GOARCH)
}
