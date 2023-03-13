package main

import (
    "fmt"
    "flag"
)

func main() {
    var (
        max int
        msg string
        x bool
    )

    flag.IntVar(&max, "n", 32, "処理数の最大値")
    flag.StringVar(&msg, "m", "", "処理メッセージ")
    flag.BoolVar(&x, "x", false, "拡張オプション")

    flag.Parse()

    fmt.Printf("処理数の最大値 = %d\n", max)
    fmt.Printf("処理メッセージ = %s\n", msg)
    fmt.Printf("拡張オプション = %v\n", x)
}
