package bubbleSort

import(
	"fmt"
)

func Bubble(numArr []int) {
	var temp int
	var swapped bool

	for i := 0; i < len(numArr)-1; i++ {
		swapped = false
		for j := 0; j < len(numArr) - i - 1; j++ {
			if numArr[j] > numArr[j+1] {
				temp = numArr[j]
				numArr[j] = numArr[j+1]
				numArr[j+1] = temp
				swapped = true
			}
		}

		if swapped == false {
			break
		}
	}

	fmt.Println(numArr)
}