package main

import (
    f "fmt"
    "os"
)

func main() {
    // os.Exit()で処理が終わるとその後に控える defer 文は破棄される
    defer func() {
        f.Println("defer print")
    }()

    f.Println("dir os main()")

    os.Exit(0)
}
