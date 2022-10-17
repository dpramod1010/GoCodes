package main

import "fmt"

type Adress struct {
	state string
	city  string
	pin   int
}
type Person struct {
	name   string
	age    int
	adress Adress
}

func add(a, b int) int {
	return a * b
}

func main() {
	fmt.Println("demo")
	p := Person{name: "pramod", age: 23,
		adress: Adress{"Maharastra", "pune", 412208},
	}
	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.adress.pin)
	nn := add(12, 12)
	fmt.Println("AA function", nn)

}

// func add(a, b, c int) int {
// 	return a * b * c
// }

// func main() {

// 	for i := 0; i <= 5; i++ {
// 		fmt.Println("..", i)
// 	}
// 	var i = 2
// 	switch i {
// 	case 1:
// 		fmt.Println("hello")
// 	case 2:
// 		fmt.Println("heoooo")

// 	}
// 	a := [4]int{23, 68, 9, 23}

// 	fmt.Println(a)

// 	m := make(map[string]int)
// 	m["pramod"] = 10
// 	m["amit"] = 112

// 	fmt.Println(m)

// 	z := add(2, 2, 2)
// 	fmt.Println("Addition is", z)
// }
