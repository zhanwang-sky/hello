package exercises

import (
	"fmt"
)

type ZeroEvenOdd struct {
	n       int
	sigZero chan bool
	sigEven chan bool
	sigOdd  chan bool
}

func (z *ZeroEvenOdd) Zero() {
	for i := range z.n {
		<-z.sigZero
		fmt.Printf("0")
		if i%2 == 0 {
			z.sigOdd <- true
		} else {
			z.sigEven <- true
		}
	}
}

func (z *ZeroEvenOdd) Even() {
	for i := range z.n / 2 {
		<-z.sigEven
		num := 2 * (i + 1)
		fmt.Printf("%d", num)
		z.sigZero <- true
	}
}

func (z *ZeroEvenOdd) Odd() {
	for i := range (z.n + 1) / 2 {
		<-z.sigOdd
		num := 2*(i+1) - 1
		fmt.Printf("%d", num)
		z.sigZero <- true
	}
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	tmp := ZeroEvenOdd{
		n:       n,
		sigZero: make(chan bool, 1),
		sigEven: make(chan bool, 1),
		sigOdd:  make(chan bool, 1),
	}
	tmp.sigZero <- true
	return &tmp
}
