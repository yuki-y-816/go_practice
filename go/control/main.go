package main

import (
    "fmt"
    con "control/constructs"
)

func main() {
    fmt.Println("--- IfFunc() ---")
    con.IfFunc()

    fmt.Println("--- ForFunc() ---")
    con.ForFunc()

    fmt.Println("--- SwitchFunc() ---")
    con.SwitchFunc()

    fmt.Println("--- GotoFunc() ---")
    con.GotoFunc()

    fmt.Println("--- GoFunc() ---")
    con.GoFunc()

    fmt.Println("--- InitFunc() ---")
    con.InitFunc()
}
