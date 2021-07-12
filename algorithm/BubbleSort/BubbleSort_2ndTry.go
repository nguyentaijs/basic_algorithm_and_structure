package main

import "fmt"

func main() {
	a := [15]int{5, 16, 456, 14523, 1463, 14, 2513, 1434, 11, 43, 59, 88, 51, 244, 166}
	fmt.Printf("Input: %v\n", a)

	k := len(a) - 1

	for k != 0 {
		n1 := 0
		for i := 0; i < k; i++ {
			if (a[i] > a[i + 1]) {
				temp := a[i]
				a[i] = a[i + 1]
				a[i + 1] = temp
				n1 = i
			}
		}
		k = n1
	}
	fmt.Printf("Output: %v\n", a)
}