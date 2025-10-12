package main

import (
	"fmt"
	"math"
	"github.com/zhanwang-sky/greetings"
)

const modName string = "hello"

const pi = 3.1415926

func main() {
	fmt.Println("modName:", modName)
	fmt.Printf("sin(π) = %.0f\n", math.Round(math.Sin(pi)))
	fmt.Println(greetings.Hello("Jon Snow"))
}
