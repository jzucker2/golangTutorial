package main

import "fmt"

func main() {
	test := [5]int{0, 1, 2, 3, 4}
	a := make([]int, 5)
	printSlice("a", a) // "a" len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	printSlice("b", b) // "b" len=0 cap=5 []
	testb := test[:0]
	printSlice("testb", testb)
	
	c := b[:2]
	printSlice("c", c) // "c" len=2 cap=5 [0 0]
	testc := testb[:2]
	printSlice("testc", testc)

	d := c[2:5]
	printSlice("c again", c)
	printSlice("d", d) // "d" len=3 cap=3 [0 0 0]
	testd := testc[2:5]
	printSlice("testd", testd)
	fmt.Printf("%v\n", test)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}