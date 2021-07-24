package main

import (
	"fmt"
)


func main() {
	s := []int{1, 4, 5, 8, 14, 36, 55, 77, 91, 100, 101, 150}
	fmt.Println(s)
	fmt.Println("Please input search val: ")
	x := getUserInt()
	left := 0
	right := len(s) - 1
	searchIdex, err := binarySearchRecursive(s, x, left, right)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Found at", searchIdex)
	}
}

func binarySearchRecursive(s []int, x int, left int, right int) (int, error) {
	if left > right {
		return 0, fmt.Errorf("Not found")
	}
	mid := (left + right) / 2
	if s[mid] == x {
		return mid, nil
	} else if s[mid] < x {
		return binarySearchRecursive(s, x, mid+1, right)
	} else {
		return binarySearchRecursive(s, x, left, mid-1)
	}
}