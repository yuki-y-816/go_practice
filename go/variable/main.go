package main

import "fmt"

var n = 100

func main() {
    n = n + 1
    a := 1
    b := 2
    // c := 3

    x, y := 1, 2

    fmt.Printf("n = %d\n", n)
    fmt.Println(a, b)
    fmt.Printf("x = %d, y = %d\n", x, y)
}
