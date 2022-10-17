package main

import "fmt"

func main() {
	fmt.Println("first For")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("Second For")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

}
