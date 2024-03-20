package main

import "fmt"

func findPositiveNumsCount() {
	var n int
	fmt.Scan(&n)

	var a []int

	for i := 0; i < n; i++ {
		var n int
		fmt.Scan(&n)
		a = append(a, n)
	}

	var c int
	for i := 0; i < n; i++ {
		if a[i] > 0 {
			c += 1
		}
	}

	fmt.Print(c)
}
