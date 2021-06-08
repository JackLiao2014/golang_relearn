package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(9))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(16))

	fmt.Printf("Now you have %g problems.\n", math.Pi)

	fmt.Println(add(42, 13))
}

func add(x int, y int) int {
	return x + y
}
