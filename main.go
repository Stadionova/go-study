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
//func main() {
//	var a int
//	fmt.Print("Введите число: ")
//	fmt.Scan(&a)
//	a *= 2
//	a += 100
//	fmt.Print(a)
//}

// 9 EXAMPLE
//func main() {
//
//	var a int
//	var b int
//	fmt.Scan(&a) // считаем переменную 'a' с консоли
//	fmt.Scan(&b) // считаем переменную 'b' с консоли
//
//	a *= a
//	b *= b
//	fmt.Println(a + b)
//}

// 10 EXAMPLE
//func main() {
//
//	var a int
//	fmt.Scan(&a)
//
//	a *= a
//	fmt.Println(a)
//}

// 11 EXAMPLE
//func main() {
//	var a int
//	fmt.Scan(&a)
//	fmt.Println(a % 10) // вывод: последняя цифра натурального числа из консоли
//}

// 12 EXAMPLE
//func main() {
//	var a int
//	fmt.Scan(&a)
//	a %= 100
//	a /= 10
//	fmt.Println(a) // вывод: вторая цифра с конца натурального числа из консоли
//}

// 13 EXAMPLE
func main() {
	var d int
	var h int
	var m int
	fmt.Scan(&d)
	h = d / 30       // округл до целого числа
	m = 2 * (d % 30) // остаток отделения умножаем на 2
	fmt.Println("It is", h, "hours", m, "minutes")
}
