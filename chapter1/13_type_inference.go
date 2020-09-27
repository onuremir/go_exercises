package main

import "fmt"

/* (either by using the := syntax or var = expression syntax),
the variable\'s type is inferred from the value on the right hand side.*/

func main() {
	v := true // bool
	fmt.Printf("v is of type %T\n", v)
	i := 42 // int
	fmt.Printf("i is of type %T\n", i)
	f := 3.142 // float64
	fmt.Printf("f is of type %T\n", f)
	g := 0.867 + 0.5i // complex128
	fmt.Printf("g is of type %T\n", g)
	s := "Hello" // string
	fmt.Printf("s is of type %T\n", s)
}
