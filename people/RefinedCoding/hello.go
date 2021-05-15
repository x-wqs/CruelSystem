package main

import "fmt"

func main() {
    var s []int
    s = append(s, 0)
    s = append(s, 1)
    s = append(s, 2, 3, 4)
    fmt.Printf("len = %d, cap = %d, %v\n", len(s), cap(s), s)
    fmt.Println("hello world")
}