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

	// Slices
	fmt.Println("Slices:")

	var s1 []int
	fmt.Printf("s1: %v, is nil? %v, len=%d\n", s1, s1 == nil, len(s1))

	s1 = make([]int, 3)
	fmt.Println("> s1 = make([]int, 3)")
	fmt.Printf("s1: %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))

	s1 = make([]int, 3, 5)
	fmt.Println("> s1 = make([]int, 3, 5)")
	fmt.Printf("s1: %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))

	s1[len(s1)-1] = -2
	s1 = append(s1, 3)
	fmt.Println("> s1[len(s1)-1] = -2")
	fmt.Println("> s1 = append(s1, 3)")
	fmt.Printf("s1: %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))

	s1 = append(s1, 4, 5)
	fmt.Println("> s1 = append(s1, 4, 5)")
	fmt.Printf("s1: %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, len(s1))
	copy(s2, s1)
	fmt.Println("s2 is the (deep) copy of s1:", s2)

	s3 := s2[2:4]
	fmt.Println("s3 is the slice (shallow copy) of s2:", s3)

	s2[2] = -s2[2]
	fmt.Println("> s2[2] = -s2[2]")
	fmt.Println("s3:", s3)

	s4 := []int{1, 2, 3, 4}
	fmt.Println("s4:", s4)

	s5 := s4[:3]
	fmt.Println("> s5 := s4[:3]")
	fmt.Println("s5:", s5)

	s5 = append(s5, 5)
	fmt.Println("> s5 = append(s5, 5)")
	fmt.Printf("s4: %v, s5: %v\n", s4, s5)

	fmt.Println()
}
