package main

import (
	"fmt"
	"math"

	"github.com/zhanwang-sky/greetings"
)

const modName string = "hello"
const pi = 3.1415926

func main() {
	// Values & Variables
	fmt.Println("Values & Variables:")

	fmt.Println("modName:", modName)
	fmt.Printf("sin(π) = %.0f\n", math.Round(math.Sin(pi)))
	fmt.Println(greetings.Hello("Jon Snow"))

	fmt.Println()

	// Arrays
	fmt.Println("Arrays:")

	var a [5]int
	fmt.Println("arr a:", a)
	fmt.Println("length of a:", len(a))

	b := [5]int{5, 4, 3, 2, 1}
	fmt.Println("arr b:", b)

	b = [...]int{2, 3, 4, 5, 6}
	fmt.Println("arr b:", b)
	b[len(b)-1] = 99
	fmt.Println("the last element of b is", b[len(b)-1])

	c := [...][3]string{
		{"a", "b", "c"},
		{"foo", "bar", "balabala"},
	}
	for i, row := range c {
		for j, val := range row {
			fmt.Printf("c[%d][%d]=%s, ", i, j, val)
		}
		fmt.Println()
	}

	fmt.Println()
}
