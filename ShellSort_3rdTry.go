package main

import "fmt"

func main() {
	a := [12]int{5, 1, 21, 456, 1234, 345, 3, 87, 623, 614, 163, 52}
	fmt.Printf("Input: %v\n", a)

	gapArr := [5]int{1,3,5,9,17}
	gap, gapIndex := getGap(a, gapArr)
	fmt.Printf("Gap: %v, Gap Index: %v\n", gap, gapIndex)

	for i := gapIndex; i >= 0; i-- {
		gap = gapArr[i]
		for j := 0; j < gapArr[gapIndex]; j++ {
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
	fmt.Printf("Output: %v", a)

}

func getGap(a [12]int, gapArr [5]int) (int, int) {
	var gap, gapIndex int
	for i := len(gapArr) - 1; i >= 0; i-- {
		if gapArr[i] < len(a)/2 {
			gap = gapArr[i]
			gapIndex = i
			break;
		}
	}
	return gap, gapIndex
}