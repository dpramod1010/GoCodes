package main

import "fmt"

func main() {
	fmt.Println("Sorting Program")
	//var temp int
	// size:=arr.length();
	//var i int;
	arr := []int{10, 2, 4, 5, 1}
	fmt.Println(arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp

			}

		}
	}
	fmt.Println("Sorted array is:-")
	for i := 0; i < len(arr); i++ {

		fmt.Println(arr[i], " ")

	}
}
