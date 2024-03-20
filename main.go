package main

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
//func main() {
//	var prevN int = 0
//	var currentN int
//	var countMaxN int = 1
//	fmt.Scan(&currentN)
//	for i := 0; currentN > 0; i++ {
//		if currentN > prevN {
//			prevN = currentN
//			countMaxN = 1
//			fmt.Scan(&currentN)
//		} else if currentN == prevN {
//			countMaxN += 1
//			fmt.Scan(&currentN)
//		} else if currentN < prevN {
//			fmt.Scan(&currentN)
//		}
//	}
//	fmt.Println(countMaxN)
//}

// 23 EXAMPLE
//func main() {
//	var n int
//	var c int
//	var d int
//	fmt.Scan(&n)
//	fmt.Scan(&c)
//	fmt.Scan(&d)
//	for i := 1; i <= n; i++ {
//		if i%c == 0 && i%d != 0 {
//			fmt.Println(i)
//			break
//		} else {
//			continue
//		}
//	}
//}

// 24 EXAMPLE
//func main() {
//	var n int
//	for i := 0; ; i++ {
//		fmt.Scan(&n)
//		if n < 10 {
//			continue
//		} else if n > 100 {
//			break
//		} else {
//			fmt.Println(n)
//		}
//	}
//}

// 25 EXAMPLE
//func main() {
//	var x int // Вклад в банке
//	var p int // Ежегодно он увеличивается на p процентов
//	var y int // максимальная сумма вклада
//	fmt.Scan(&x)
//	fmt.Scan(&p)
//	fmt.Scan(&y)
//	var sumWithPercent int = x
//	for i := 0; ; i++ {
//		if sumWithPercent < y {
//			var percentFromSum int = (sumWithPercent * p) / 100 // процент от x суммы
//			sumWithPercent += percentFromSum / 1                // отбрасываем дробную часть
//		} else {
//			fmt.Println(i)
//			break
//		}
//	}
//}

// 26 EXAMPLE
//func main() {
//	var n int
//	var d int
//
//	fmt.Scan(&n)
//	fmt.Scan(&d)
//
//	var nCopy int = n // 564
//	var dCopy int = d // 8954
//
//	for i := 0; i <= 5; i++ {
//		var n1 = nCopy/10 == 0    // 1 digit 1
//		var n2 = nCopy/100 == 0   // 2 digit 12
//		var n3 = nCopy/1000 == 0  // 3 digit 123
//		var n4 = nCopy/10000 == 0 // 4 digit 1234
//
//		var nToFind int = nCopy // first digit n
//		if n1 {
//			nToFind = nCopy
//		} else if n2 {
//			nToFind = nCopy / 10
//		} else if n3 {
//			nToFind = nCopy / 100
//		} else if n4 {
//			nToFind = nCopy / 1000
//		}
//		dCopy = d
//
//		for j := 0; j <= 5; j++ {
//			var d1 = dCopy/10 == 0    // 1 digit 1
//			var d2 = dCopy/100 == 0   // 2 digit 12
//			var d3 = dCopy/1000 == 0  // 3 digit 123
//			var d4 = dCopy/10000 == 0 // 4 digit 1234
//			var dToFind = dCopy       // first digit d
//
//			if d1 {
//				dToFind = dCopy
//			} else if d2 {
//				dToFind = dCopy / 10
//			} else if d3 {
//				dToFind = dCopy / 100
//			} else if d4 {
//				dToFind = dCopy / 1000
//			}
//
//			if d1 {
//				dCopy = dCopy
//			} else if d2 {
//				dCopy = dCopy % 10 // cut first digit d
//			} else if d3 {
//				dCopy = dCopy % 100 // cut first digit d
//			} else if d4 {
//				dCopy = dCopy % 1000 // cut first digit d
//			}
//
//			if nToFind == dToFind {
//				fmt.Print(nToFind)
//				fmt.Print(" ")
//				break
//			}
//		}
//
//		if n1 {
//			break
//		} else if n2 {
//			nCopy = nCopy % 10 // cut first digit n
//		} else if n3 {
//			nCopy = nCopy % 100 // cut first digit n
//		} else if n4 {
//			nCopy = nCopy % 1000 // cut first digit n
//		}
//	}
//}

// 27 EXAMPLE
//func main() {
//	formatData()
//}

// 28 EXAMPLE
//func main() {
//	workArray()
//}

// 29 EXAMPLE
//func main() {
//	var n int // n >= 4
//	fmt.Scan(&n)
//	var a []int
//
//	for i := 0; i < n; i++ {
//		var elem int
//		fmt.Scan(&elem)
//		a = append(a, elem)
//	}
//	fmt.Print(a[3])
//}

// 30 EXAMPLE
func main() {
	findMaxNum()
}
