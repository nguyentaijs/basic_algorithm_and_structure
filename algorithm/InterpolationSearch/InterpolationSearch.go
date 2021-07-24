package main

import (
    "bufio"
    "errors"
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
    result, err := interpolationSearch(s, left, right, x)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Found", result)
    }
}

func interpolationSearch(s []int, left int, right int, x int) (int, error) {

    for left <= right && s[left] <= x && s[right] >= x {

        mid := left + (x - s[left]) * ((right - left) / (s[right] - s[left]))
        if s[mid] == x {
            return mid, nil
        }
        if s[mid] < x {
            return interpolationSearch(s, mid + 1, right, x)
        }
        if s[mid] > x {
            return interpolationSearch(s, left, mid - 1, x)
        }
    }
    return 0, errors.New("not found")
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
