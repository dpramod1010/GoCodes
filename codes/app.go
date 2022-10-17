package main

import "fmt"

func main() {
	// var conferenceName = "Go Conference"
	// const totalTicket = 50
	// var remainTicket = 50

	// fmt.Printf("Welcome to %v booking application\n", conferenceName)
	// fmt.Printf("We Have Total %v Tickets And %v are Remaining  \n", totalTicket, remainTicket)

	// var name = "pramod"
	// fmt.Printf("%v", name)

	// var nn string
	// fmt.Println("Enter A Student Name")
	// fmt.Scan(&nn)
	// fmt.Println("Name Is:", nn)

	// var name string

	// fmt.Println("Enter User Name:-")
	// fmt.Scan(&name)
	// var Number string
	// if len(name) > 2 {
	// 	fmt.Println("Enter Contact Number:-")
	// 	fmt.Scan(&Number)
	// 	fmt.Println("len of int", len(Number))
	// 	if len(Number) < 10 || len(Number) > 10 {
	// 		fmt.Println("Invalid Number Plz Enter Correct Number:")
	// 	}
	// } else {
	// 	fmt.Println("Name Should Be Grater Than 2 Chareter")
	// }

	// // /*var Number string
	// // fmt.Println("Enter Contact Number:-")
	// // fmt.Scan(&Number)
	// // fmt.Println("len of int", len(Number))
	// // if len(Number) < 10 || len(Number) > 10 {
	// // 	fmt.Println("Invalid Number Plz Enter Correct Number:")
	// // }*/

	// var uChoice int
	// for {

	// 	fmt.Println("Press 1 For Book Tickets :- ")
	// 	fmt.Println("Press 2 For Exit :-  ")
	// 	fmt.Scan(&uChoice)

	// 	if uChoice == 1 {

	// 		const totalTick = 50
	// 		var bookTick int32
	// 		var remTic int32

	// 		fmt.Println("Enter How Many Tic Wants To Book: - ")
	// 		fmt.Scan(&bookTick)

	// 		fmt.Printf("Thank You %v You Have Booked Total %v Tickets :-   \n", name, bookTick)
	// 		remTic = totalTick - bookTick
	// 		fmt.Printf("SMS Sucessfully Send to %v \n", Number)
	// 		fmt.Printf("Total Remaining  Tickets Are :- %v\n ", remTic)
	// 		//fmt.Printf("Your Contact Number is:- %v \n", cNo)

	// 	} else {
	// 		fmt.Println("Thanks....")
	// 		break
	// 	}
	// }

	var new []int
	var n int
	// var input int
	// fmt.Println("How many Input")
	// fmt.Scan(&input)
	var nn int
	fmt.Println("Enter How many")
	fmt.Scan(&nn)
	for i := 1; i <= nn; i++ {
		fmt.Println("Enter Value To Add in Sclice:-  ")
		fmt.Scan(&n)
		new = append(new, n)

		fmt.Printf("%v is Added in sclice , Enter New number to add :-\n ", n)
	}
	fmt.Println("Total Added Sclice Is:- ", new)
	fmt.Println("Enter Index Number to get Value find:- ")
	var f int
	fmt.Scan(&f)

	fmt.Printf("%v Is in %v th Position ", n, f)

	// beforeRemoveSlice := []int{1, 2, 3, 4, 5}
	// indexRemove := 2

	var indexRemove int
	var p int
	fmt.Println("How many Numbers u Wants to delete:")
	fmt.Scan(&p)
	for i := 1; i <= p; i++ {

		fmt.Println("Enter Index to remove from Sclice:-  ")
		fmt.Scan(&indexRemove)

		fmt.Println("before remove:", new)

		afterRemoveSlice := remove(new, indexRemove)
		fmt.Println("after remove", afterRemoveSlice)

	}
}
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

// ........................	var remTic int32
// 	var bookTick int32

// 	fmt.Println("Enter How Many Tic Wants To Book: - ")
// 	fmt.Scan(&bookTick)
// 	fmt.Printf("Thank You %v Have Booked Total %v Tickets\n", uname, bookTick)
// 	remTic = totalTick - bookTick

// 	fmt.Printf("Total Remain Tickets Are :- %v\n ", remTic)

// 	if remTic < 1 {
// 		fmt.Printf("%v Tickets Are Remain Please Try Again", remTic)
// 		fmt.Println("HouseFulll..............|")
// 	}
// }

// //var array []string
// // array[0] = "Pramod"
// // array[1] = "Dhayarkar"
// // array[10] = "10 Array"

// var slice []int

// slice = append(slice, 10, 12, 13)

// slice = append(slice, 14)
// slice = append(slice, 89)

// fmt.Printf("Total Array is:- %v\n", slice)

// fmt.Printf("First Positon Of an array is:- %v\n", slice[3])
// fmt.Printf("Type of an array is :- %T\n", slice)

// fmt.Printf("Length of an array is :%v", len(slice))
