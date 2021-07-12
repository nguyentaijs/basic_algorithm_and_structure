package main

import "fmt"

func main() {
	a := [10]int{4524, 76, 1, 743, 99, 3, 65, 132, 6, 334}
	fmt.Printf("Input: %v\n", a)
	bubbleSort(&a)
	fmt.Printf("Output: %v", a)

}

func bubbleSort(a *[10]int) {
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
}