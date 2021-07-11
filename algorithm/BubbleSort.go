package main

import "fmt"

func main() {
	a := [15]int{5, 613, 75456, 16, 672, 146, 6, 712, 21, 74, 252, 6777, 33, 133, 5656}
	fmt.Printf("Input: %v\n", a)

	k := len(a) - 1
	for k != 0 {
		n1 := 0
		for i := 0; i < k; i++ {
			if a[i] > a[i + 1] {
				temp := a[i]
				a[i] = a[i + 1]
				a[i + 1] = temp
				n1 = i
			}
		}
		k = n1
	}
	fmt.Printf("Output: %v", a)

}