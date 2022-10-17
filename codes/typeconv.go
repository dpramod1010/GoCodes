package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Using Atoi() function
	a := "244"
	c := "12"
	b, e := strconv.Atoi(a)
	d, e := strconv.Atoi(c)

	if e == nil {
		fmt.Printf("%T \n %v", b, b)
		fmt.Printf("%T \n %v", d, d)
	}

}

// func main() {
// 	z := "12"
// 	z, e := strconv.Atoi(z)
// 	if e == nil {
// 		fmt.Printf("%T \n %v", z, z)
// 	}
// }
