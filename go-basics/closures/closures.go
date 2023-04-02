package main

import "fmt"

func sequence() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	nextSeq := sequence()

	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
}
