package binarySearch

import(
	"fmt"
)

func BiSearch(numArr []int, find int) {
	var l int = 0
	var r int = len(numArr)-1
	var res int = 0

	for l <= r {
		m := (l+r)/2
		if numArr[m] == find { 
			res = m
			break
		} else if numArr[m] < find {
			l = m+1
		} else if numArr[m] > find {
			r = m-1
		} else {
			res = -1
		}
	}

	if res > 0 {
		fmt.Println("The number can be found at array index[",res,"]")
	} else {
		fmt.Println("The number does not appear to be in the array")
	}
}