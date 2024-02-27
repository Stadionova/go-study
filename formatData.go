package main

import "fmt"

func formatData() {
	var n float64
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Printf("число %2.2f не подходит", n)
	} else if n > 10000 {
		fmt.Printf("%e", n)
	} else {
		fmt.Printf("%.4f", n*n)
	}
}
