package main

import (
	"fmt"
	"maps"
	"math"
	"slices"

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
	if slices.Equal(s1, s2) {
		fmt.Println("s1==s2")
	} else {
		fmt.Println("s1!=s2")
	}

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

	// Maps
	fmt.Println("Maps:")

	var m map[string]int
	fmt.Println("> var m map[string]int")
	fmt.Printf("m: %v, m is nil? %v\n", m, m == nil)

	m = make(map[string]int)
	fmt.Println("> m = make(map[string]int)")
	fmt.Printf("m: %v, m is nil? %v\n", m, m == nil)

	m["1"] = '1'
	m["a"] = 'a'
	fmt.Printf("m: %v, len(m)=%d\n", m, len(m))

	delete(m, "a")
	fmt.Println("> delete(m, \"a\")")
	fmt.Printf("m: %v, len(m)=%d\n", m, len(m))
	delete(m, "b")
	fmt.Println("> delete(m, \"b\")")
	fmt.Printf("m: %v, len(m)=%d\n", m, len(m))
	clear(m)
	fmt.Println("> clear(m)")
	fmt.Printf("m: %v, len(m)=%d\n", m, len(m))

	_, prs := m["foo"]
	fmt.Println("is \"prs\" exists?", prs)
	fmt.Println("map[\"key\"] will not create \"key\" automatically!")
	fmt.Printf("m: %v, len(m)=%d\n", m, len(m))

	m1 := map[string]int{"foo": 1, "bar": 2}
	m2 := map[string]int{"bar": 2, "foo": 1}
	fmt.Printf("m1: %v, m2: %v\n", m1, m2)
	if maps.Equal(m1, m2) {
		fmt.Println("m1==m2")
	} else {
		fmt.Println("m1!=m2")
	}

	fmt.Println()
}
