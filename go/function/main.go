package main

import (
    "fmt"
)

func plus(x, y int) int {
    return x + y
}

func hello() {
    fmt.Printf("Hello, World.\n")
}

func div(a, b int) (int, int) {
    q := a / b
    r := a % b

    return q, r
}

var plusAlias = plus

func returnFunc() func() {
    return func() {
        fmt.Printf("I'm a function\n")
    }
}

func callFunc(f func()) {
    f()
}

func later() func(string) string {
    var store string

    return func(next string) string {
        s := store
        store = next

        return s
    }
}

func integers() func() int {
    i := 1

    return func() int {
        i++

        return i
    }
}

func main() {
    fmt.Printf("7 + 10 = %v\n", plus(7, 10))

    hello()

    q, r := div(19, 7)
    fmt.Printf("19 / 7 = %v\n", q)
    fmt.Printf("19 %% 7 = %v\n", r)
    q1, _ := div(19, 7)
    fmt.Printf("q1 = %v\n", q1)

    f := func(x, y int) int { return x + y}
    fmt.Printf("f(7, 8) = %v\n", f(7, 8))
    fmt.Printf("f() = %T\n", f)
    fmt.Printf("%v\n", func(x,y int) int {return x + y} (2, 3))

    fmt.Printf("plusAlias(2, 3) = %v\n", plusAlias(2, 3))

    f1 := returnFunc()
    f1()
    returnFunc()()

    callFunc(hello)
    callFunc(func(){
        fmt.Printf("this is callFunc()\n")
    })

    f2 := later()
    fmt.Println(f2("Golang"))
    fmt.Println(f2("Golang2"))
    fmt.Println(f2("Golang3"))

    i1 := integers()
    fmt.Println(i1())
    fmt.Println(i1())
    fmt.Println(i1())

    i2 := integers()
    fmt.Println(i2())
    fmt.Println(i2())
}
