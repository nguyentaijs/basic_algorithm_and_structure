package main

import "fmt"

func main() {
	var inputArr = [12]int{10, 57, 20, 999, 35, 14, 65, 116, 189, 822, 1024, 1}
	fmt.Printf("Array before sorting: %v\n", inputArr)

	// Input array
	var papenovGapList = [7]int{1, 3, 5, 9, 17, 33, 65}
	
	var gap int
	var gapIndex int

	// Get gap < inputArray.length / 2
	for i := len(papenovGapList) - 1; i >= 0; i-- {
		curGap := papenovGapList[i]
		if ( curGap < len(inputArr) / 2) {
			gapIndex = i
			break;
		}
	}
	gap = papenovGapList[gapIndex]
	fmt.Printf("Gap used: %v\n", gap)

	for i := gapIndex; i >= 0; i-- {
		gap = papenovGapList[i]
		fmt.Printf("----------- Interation gap = %v -----------\n", gap)
		for dc := 0; dc < gap; dc++ {
			fmt.Printf("Sort the sub array subArr = %v by insertion sort\n", dc);
			for i := dc; i < len(inputArr); i = i + gap {
				fmt.Printf("a[%v] = %v, ", i, inputArr[i])
			}
			fmt.Printf("\n")
			// Insertion sort
			for i := dc + gap; i < len(inputArr); i += gap {
				fmt.Printf("Consider the element: %v\n", i)
				t := inputArr[i]
				j := i

				for j > dc && inputArr[j - gap] > t {
					inputArr[j] = inputArr[j - gap]
					j -= gap
				}
				if (j < i) {
					inputArr[j] = t
				}
			}
			for k := dc; k < len(inputArr); k = k + gap {
				fmt.Printf("a[%v] = %v, ", k, inputArr[k])
			}
			fmt.Printf("\n")
		}
		
	}

}