package main

import ( //factored import statement
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random number is", rand.Intn(100))
}
