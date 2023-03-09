package main

import (
    "fmt"
    "os"
    "log"
)

func main() {
    _, err := os.Open("foobar")
    if err != nil {
        log.Fatal(err)
    }
}
