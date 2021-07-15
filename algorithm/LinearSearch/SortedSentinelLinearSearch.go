package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := []int{1, 2, 4, 22, 25, 45, 50, 100, 111, 423, 644, 666}
	log.Printf("Input: %v\nPlease input search int:", s)

	searchVal := getUserInt()
	// Append sentinel to slice
	s = append(s, searchVal)
	i, err := sentinelLinearSearch(s, searchVal)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Found at", i)
	}
}

func sentinelLinearSearch(s []int, searchVal int) (int, error) {
	i := 0
	for s[i] < searchVal {
		i++
	}
	if (s[i] == searchVal) {
		return i, nil
	}
	return i, fmt.Errorf("Not found")
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