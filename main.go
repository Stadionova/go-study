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
//func main() {
//	var d int
//	var h int
//	var m int
//	fmt.Scan(&d)
//	h = d / 30       // округл до целого числа
//	m = 2 * (d % 30) // остаток отделения умножаем на 2
//	fmt.Println("It is", h, "hours", m, "minutes")
//}

// 14 EXAMPLE
//func main() {
//	var d int
//	fmt.Scan(&d)
//
//	if d > 0 {
//		fmt.Println("Число положительное")
//	} else if d < 0 {
//		fmt.Println("Число отрицательное")
//	} else {
//		fmt.Println("Ноль")
//	}
//}

// 15 EXAMPLE
//func main() {
//	var d int
//	fmt.Scan(&d)
//	var n1 int = d % 10
//	d = d / 10
//	//fmt.Println(n1)
//	var n2 int = d % 10
//	d = d / 10
//	//fmt.Println(n2)
//	var n3 int = d % 10
//	d = d / 10
//	//fmt.Println(n3)
//	if n1 != n2 && n1 != n3 && n2 != n3 {
//		fmt.Println("YES")
//	} else {
//		fmt.Println("NO")
//	}
//}

// 16 EXAMPLE
//func main() {
//	var d int
//
//	var oneD int = 0      // 1
//	var twoD int = 10     // 12
//	var threeD int = 100  // 123
//	var fourD int = 1000  // 1234
//	var fiveD int = 10000 // 12345
//	var maxD int = 100000 // 123456
//
//	fmt.Scan(&d)
//
//	if d >= oneD && d < twoD {
//		fmt.Println(d % 10)
//	} else if d >= twoD && d < threeD {
//		fmt.Println(d / 10)
//	} else if d >= threeD && d < fourD {
//		fmt.Println(d / 100)
//	} else if d >= fourD && d < fiveD {
//		fmt.Println(d / 1000)
//	} else if d >= fiveD && d < maxD {
//		fmt.Println(d / 10000)
//	}
//}

// 17 EXAMPLE
//func main() {
//	var luckyD int
//	fmt.Scan(&luckyD)
//
//	var firstThreeD int = luckyD / 1000
//	var fD int = firstThreeD / 100
//	var sD int = firstThreeD / 10 % 10
//	var tD int = firstThreeD % 10
//	var firstThreeSum int = fD + sD + tD
//
//	var secondThreeD int = luckyD % 1000
//	var fSD int = secondThreeD / 100
//	var sSD int = secondThreeD / 10 % 10
//	var tSD int = secondThreeD % 10
//	var secondThreeSum int = fSD + sSD + tSD
//
//	if firstThreeSum == secondThreeSum {
//		fmt.Println("YES")
//	} else {
//		fmt.Println("NO")
//	}
//}

// 18 EXAMPLE
//func main() {
//	var d int
//	fmt.Scan(&d)
//
//	var firstCondition int = d % 400
//	var secondCondition int = d % 4
//	var thirdCondition int = d % 100
//
//	if firstCondition == 0 || (secondCondition == 0 && thirdCondition != 0) {
//		fmt.Println("YES")
//	} else {
//		fmt.Println("NO")
//	}
//}

// 19 EXAMPLE
//func main() {
//	for i := 1; i <= 10; i++ {
//		fmt.Println(i * i)
//	}
//}

// 20 EXAMPLE
//func main() {
//	var A int
//	var B int
//	var sumAB int = 0
//
//	fmt.Scan(&A)
//	fmt.Scan(&B)
//
//	for i := A; i <= B; i++ {
//		sumAB += i
//	}
//	fmt.Println(sumAB)
//}

// 21 EXAMPLE
//func main() {
//	var countN int
//	var eachN int
//	var sumN int
//
//	fmt.Scan(&countN)
//	fmt.Println(sumN)
//	for i := 1; i <= countN; i++ {
//		fmt.Scan(&eachN)
//		if eachN%8 == 0 && eachN > 9 && eachN < 100 {
//			sumN += eachN
//		}
//	}
//
//	fmt.Println(sumN)
//}

// 22 EXAMPLE
func main() {
	var prevN int = 0
	var currentN int
	var countMaxN int = 1
	fmt.Scan(&currentN)
	for i := 0; currentN > 0; i++ {
		if currentN > prevN {
			prevN = currentN
			countMaxN = 1
			fmt.Scan(&currentN)
		} else if currentN == prevN {
			countMaxN += 1
			fmt.Scan(&currentN)
		} else if currentN < prevN {
			fmt.Scan(&currentN)
		}
	}
	fmt.Println(countMaxN)
}
