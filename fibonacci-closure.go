package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	p1, p2, p3 := 0, 1, 1
	count := 0
	return func() int {
		count++
		if count == 1 {
			return p1
		} else if count == 2 {
			return p2
		} else {
			p3 = p1 + p2
			p1, p2 = p2, p3
			return p3
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
