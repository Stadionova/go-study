package main

import "fmt"

func workArray() {
	var workArray [10]uint8 // 1 2 3 4 5 6 7 8 9 10
	for j := 0; j < 10; j++ {
		var n uint8
		fmt.Scan(&n)
		workArray[j] = n
	}

	for j := 0; j < 3; j++ {
		var indexFirst int  // 0 index
		var indexSecond int // 1 index
		fmt.Scan(&indexFirst)
		fmt.Scan(&indexSecond)
		var firstArrN = workArray[indexFirst]   // 1
		var secondArrN = workArray[indexSecond] // 2
		workArray[indexFirst] = secondArrN      // 2
		workArray[indexSecond] = firstArrN      // 1
	}

	for j := 0; j < 10; j++ {
		fmt.Print(workArray[j])
		fmt.Print(" ")
	}
}
