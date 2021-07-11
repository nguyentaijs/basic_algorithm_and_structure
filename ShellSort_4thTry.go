package main

import "fmt"

func main() {
	a := [20]int{5, 1, 34, 84, 127, 27, 3248, 15, 77, 12345, 85, 346456, 2432, 123, 2, 65, 3, 0, 22, 413}
	fmt.Printf("Input: %v\n", a)

	gapArray := [5]int {1,3,5,9,17}
	var gapIndex int
	for i := len(gapArray) - 1; i >= 0; i-- {
		if gapArray[i] < len(a) / 2 {
			gapIndex = i
			break
		}
	}
	fmt.Printf("Gap Index: %v\n", gapIndex)

	for i := gapIndex; i >= 0; i-- {
		gap := gapArray[i]
		for j := 0; j < gap; j++ {
			for k := j + gap; k < len(a); k += gap {
				curVal := a[k]
				curPos := k

				for curPos > j && a[curPos - gap] > curVal {
					a[curPos] = a[curPos - gap]
					curPos -= gap
				}
				if (curPos < k) {
					a[curPos] = curVal
				}
			}
			
		}
	}
	fmt.Printf("Output: %v", a)
}