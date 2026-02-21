// packages in go
package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("My favorite number is ", rand.Intn(12))

	// exported names have a capital letter
	fmt.Println(math.Pi)
	fmt.Println(add(23, 56))
}

// functions in go
func add(x int, y int) int {
	return x + y
}
