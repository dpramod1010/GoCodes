package main

import "fmt"

func Add(a, b int) int {
	return a * b
}

func main() {
	Add(3, 5)
	fmt.Println(Add)
	new := []int{10, 22, 19}
	fmt.Println(new)
}
