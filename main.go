package main

import (
	"fmt"
	"math"
)

const modName string = "hello"

const pi = 3.1415926

func main() {
	fmt.Println("modName:", modName)
	fmt.Printf("sin(π) = %.0f\n", math.Round(math.Sin(pi)))
}
