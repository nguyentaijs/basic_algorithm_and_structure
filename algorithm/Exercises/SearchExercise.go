package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strings"
)

func main() {
    s := []string{"Test", "An", "Binh", "Cong", "Dung", "Em", "Gia", "Hieu", "Ien", "Jan", "Kien"}
    fmt.Println(s)
    fmt.Println("Please input search string")
    searchString, err := getUserInput()
    if err != nil {
        panic(err)
    }

    index, err := linearSearch(s, searchString)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Found at", index)

}

func getUserInput() (string, error) {
    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if err != nil {
        return "", err
    }
    input = strings.ReplaceAll(input, "\r\n", "")
    return input, nil
}

func linearSearch(s []string,searchString string) (int, error) {
    for i := 0; i < len(s); i++ {
        if strings.EqualFold(s[i], searchString) {
            return i, nil
        }
    }
    return 0, errors.New("Not found")

}