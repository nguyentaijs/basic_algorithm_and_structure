package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
	s := []int{1, 4, 5, 8, 14, 36, 55, 77, 91, 100, 101, 150}
	fmt.Println(s)
	fmt.Println("Please input search val: ")
	x := getUserInt()
	left := 0
	right := len(s) - 1
	searchIdex, err := binarySearchIterate(s, x, left, right)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Found at", searchIdex)
	}
}

func binarySearchIterate(s []int, x int, left int, right int) (int, error) {
	for left < right {
		mid := (left + right) / 2
		if s[mid] == x {
			return mid, nil
		} else if s[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0, fmt.Errorf("Not found")
}

func getUserInt() int {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	text = strings.ReplaceAll(text, "\r\n", "")
	result, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return result
}