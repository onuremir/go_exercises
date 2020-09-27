package main

import "fmt"

const Pi = 3.14

func main() {
	const Place = "Istanbul"
	fmt.Println("Where am i?", Place)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
