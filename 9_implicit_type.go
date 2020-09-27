package main

import "fmt"

//x := 10 //error just for function body "non-declaration statement outside function body"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
