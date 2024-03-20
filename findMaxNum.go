package main

import "fmt"

func findMaxNum() {
	var a []int

	for i := 0; i < 5; i++ {
		var n int
		fmt.Scan(&n)
		a = append(a, n)
	}

	var j int = 1
	var x int = 0
	for i := 0; i < 4; i++ {
		if a[x] > a[j] {
			j += 1
		} else if a[x] <= a[j] {
			a[x] = a[j]
			a[j] = a[x]
			j += 1
		}
	}
	fmt.Print(a[0])
}
