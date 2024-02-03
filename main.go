package main

import "fmt"

// 1 EXAMPLE
// func main() {
//     fmt.Println("Hello, Go!")
// }

// 2 EXAMPLE
// func main() {

//     var hello string

//     hello = "Hello Go!"

//     var a int = 2019

//     fmt.Println(hello)
//     fmt.Println(a)

// 	var symbol int32 = 'c'
// 	fmt.Println(string(symbol))

// }

// 3 EXAMPLE
//func main() {
//	var (
//		name string = "Dima"
//		age  int    = 23
//	)
//
//	fmt.Println(name)
//	fmt.Println(age)
//}

// 4 EXAMPLE
//func main() {
//	var name string
//	var age int
//	fmt.Print("Введите имя: ")
//	fmt.Scan(&name)
//	fmt.Print("Введите возраст: ")
//	fmt.Scan(&age)
//
//	fmt.Println(name, age)
//}

// 5 EXAMPLE
//func main() {
//	var name string
//	var age int
//	fmt.Print("Введите имя и возраст: ")
//	fmt.Scan(&name, &age)
//	fmt.Println(name, age)
//}

// 6 EXAMPLE
//func main() {
//	//fmt.Print("hello, world")
//	//fmt.Print("hello, world")
//	fmt.Println("hello, world")
//	fmt.Print("hello, world")
//}

// 7 EXAMPLE
//func main() {
//	name := "Ivan"
//	age := 27
//	fmt.Println("My name is", name, "and I am", age, "years old.")
//}

// 8 EXAMPLE
func main() {
	var a int
	fmt.Print("Введите число: ")
	fmt.Scan(&a)
	a *= 2
	a += 100
	fmt.Print(a)
}
