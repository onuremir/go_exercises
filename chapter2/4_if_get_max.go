package main

import "fmt"

func main() {
	fmt.Println(getMax(20, 170))
}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
