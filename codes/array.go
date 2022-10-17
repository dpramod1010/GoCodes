package main

import (
	"fmt"
	"sort"
)

func main() {
	var a [5]int //initialized an array not assign value
	fmt.Println(a)

	a[2] = 10
	fmt.Println(a)

	b := [4]int{50, 20, 5, 10}
	fmt.Println(b)

	//sort array in ascending order
	c := b[:]
	sort.Ints(c)
	fmt.Println("Ascending  Array Is:", c)

	//sort array in descending
	d := b[:]
	sort.Sort(sort.Reverse(sort.IntSlice(d)))
	fmt.Println("Desc Array is:", d)

}
