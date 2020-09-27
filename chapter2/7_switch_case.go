package main

import (
	"fmt"
	"runtime"
)

/*
cases NO need break statement
 switch cases need not be constants, and the values involved need not be integers
*/
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
