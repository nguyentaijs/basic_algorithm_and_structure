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
	s := []int {1,4,34,2,56,29,85,50,421,111,423,64,666}
	log.Printf("Input: %v\nPlease input search int:", s)

	searchVal := getUserInt()
	i, err := linearSearch(s, searchVal)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Found at", i)
	}
}

func linearSearch(s []int, searchVal int) (int, error) {
	for i := 0; i < len(s); i++ {
		if s[i] == searchVal {
			return i, nil
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
