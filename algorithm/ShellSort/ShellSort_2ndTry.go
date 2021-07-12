package main

import "fmt"

func main() {
	a := [12]int{31, 4, 20, 5, 7, 234, 53, 111, 12, 643, 2134, 64}
	fmt.Printf("Input: %v\n", a)

	gapArray := [7]int{1,3,5,9,17,33,65}
	fmt.Printf("Gap array: %v\n", gapArray)

	gap, gapIndex := getGap(a, gapArray)
	fmt.Printf("GapIdx: %v, gap: %v\n", gapIndex, gap)
	for i := gapIndex; i >= 0; i-- {
		gap = gapArray[i]
		for j := 0; j < gap; j++ {
			for k := j + gap; k < len(a); k += gap {
				curPos := k
				curVal := a[k]

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
	fmt.Printf("Output: %v\n", a)


}

func getGap(a [12]int, gapArray [7]int) (int, int) {
	var gap int
	var gapIndex int

	for i := len(gapArray) - 1; i >= 0; i-- {
		if gapArray[i] <= len(a)/2 {
			gap = gapArray[i]
			gapIndex = i
			break
		}
	}
	return gap, gapIndex
}