package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a) //[0，0，0，0，0];5,5

	b := make([]int, 0, 5)
	printSlice("b", b) //[],0,5

	c := b[:2]
	printSlice("c", c) //[],2,2

	d := c[2:5]
	printSlice("d", d)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(s), cap(x), x)
}
