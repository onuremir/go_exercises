package main

import "fmt"

func main() {
	res := add(5, 78)
	fmt.Println("result is ", res)
}

func add(v1 int, v2 int) int {
	return v1 + v2
}
