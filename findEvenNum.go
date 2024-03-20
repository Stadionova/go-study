package main

import "fmt"

func findEvenNum() {
	var n int
	fmt.Scan(&n)

	var a []int

	for i := 0; i < n; i++ {
		var n int
		fmt.Scan(&n)
		a = append(a, n)
	}

	var b []int
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			b = append(b, a[i])
		}
	}

	for i := 0; i < len(b); i++ {
		fmt.Print(b[i])
		fmt.Print(" ")
	}
}
