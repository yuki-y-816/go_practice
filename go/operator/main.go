package main

import (
    "fmt"
)

func main () {
    fmt.Printf("4 + 2 = %v\n", 4 + 2)
    fmt.Printf("41 - 7 = %v\n", 41 - 7)
    fmt.Printf("3 * 7 = %v\n", 3 * 7)
    fmt.Printf("12 / 4 = %v\n", 12 / 4)
    fmt.Printf("5,2 の剰余 = %v\n", 5 % 2)

    s := "Taro" + " " + "Yamada"
    fmt.Printf("%v\n", s)

    fmt.Printf("AND : %v\n", 165 & 155)
    fmt.Printf("OR : %v\n", 197 | 237)
    fmt.Printf("XOR : %v\n", 92 ^ 137)
    fmt.Printf("AND NOT : %v\n", 108 &^ 13)
    fmt.Printf("<< : %v\n", 1 << 10)
    fmt.Printf(">> : %v\n", 600 >> 2)

    fmt.Printf("1 == 1 : %v\n", 1 == 1)
    fmt.Printf("true != !false : %v\n", true != !false)
    fmt.Printf("true && true : %v\n", true && true)
    fmt.Printf("true || false : %v\n", true || false)
}
