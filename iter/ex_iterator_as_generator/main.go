package main

import (
	"fmt"
	. "github.com/LearnLoop365/go-lang-util/iter"
)

// Fibo _ Fibonacci Sequence Generator
// *Fibo implements Iterable[uint], Iterator[uint]
type Fibo struct {
	val  uint
	prev uint

	max uint
}

// *Fibo implementing Iterator[uint]

func (fb *Fibo) HasNext() bool {
	return fb.prev+fb.val <= fb.max
}

func (fb *Fibo) Next() uint {
	var tmp uint = fb.prev
	fb.prev = fb.val
	fb.val += tmp
	return fb.val
}

// Ensure *Fibo Implements Iterator[uint]
var _ Iterator[uint] = (*Fibo)(nil)

func NewFibo(max uint) *Fibo {
	return &Fibo{val: 0, prev: 1, max: max}
}

//  *Fibo implementing Iterable[uint]

func (fb *Fibo) GetIterator() Iterator[uint] {
	return fb
}

// Ensure *Fibo Implements Iterable[Int]
var _ Iterable[uint] = (*Fibo)(nil)

// Sum of Even Fibonacci Numbers Not Exceeding 4,000,000
func main() {
	var sum uint = 0
	var fibo Iterable[uint] = NewFibo(4_000_000)

	// Iteration over an Iterable
	ForIn(fibo, func(it uint) {
		if it%2 == 0 {
			sum += it
		}
	})

	fmt.Println(sum)
}
