package main

import (
	"fmt"
)

type counter uint

func (c counter) get() uint {
	return uint(c)
}

func (c counter) increment() { // BUG Pointer receiver
	// fmt.Println("increment(): Before incrementing:", c)
	c = +1 // BUG it was ment to be *c+=1 pt *c++
	// fmt.Println("increment(): After incrementing:", c)
}

func main() {
	c := counter(10)
	c.increment()
	// expected c == 11
	fmt.Println("final value:", c.get())
}
