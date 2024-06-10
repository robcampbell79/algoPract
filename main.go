package main

import(
	"fmt"
	"math/rand/v2"

	"algoPract/binarySearch"
	"algoPract/bubbleSort"
)

func main() {
	var option int
	var numArr []int

	fmt.Printf("Please select an algorithm\n1: Binary Search\n2: Bubble Sort\n")

	fmt.Scanln(&option)

	eternalLoop:
	for {
		switch(option) {
		case 1:
			ranNum := rand.IntN(50)

			for i := range 51 {
				numArr = append(numArr, i+2)
			}
		
			fmt.Println(numArr)
			fmt.Println("The number to find is: ", ranNum)
			binarySearch.BiSearch(numArr, ranNum)
			break eternalLoop
		case 2:
			for i := range 11 {
				numArr = append(numArr, rand.IntN(50))
				_ = i
			}
			fmt.Println(numArr)
			bubbleSort.Bubble(numArr)
			break eternalLoop
		default:
			break eternalLoop
		}
	}

}