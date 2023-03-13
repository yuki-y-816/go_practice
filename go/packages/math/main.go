package main

import (
    "fmt"
    "math"
)

func main(){
    fmt.Println("--- 定数 ---")
    fmt.Println(math.Pi)
    fmt.Println(math.Phi)
    fmt.Println(math.E)
    fmt.Println(math.Sqrt2)
    fmt.Println(math.MaxFloat64)
    fmt.Println(math.SmallestNonzeroFloat64)
    fmt.Println(math.MaxInt64)

    fmt.Println("--- 絶対値 ---")
    fmt.Println(math.Abs(-3.14))

    fmt.Println("--- 累乗 ---")
    fmt.Println(math.Pow(2, 10))
    fmt.Println(math.Pow(2, 0))

    fmt.Println("--- 平方根と立方根 ---")
    fmt.Println(math.Sqrt(2))
    fmt.Println(math.Cbrt(8))

    fmt.Println("--- 最大値と最小値 ---")
    fmt.Println(math.Max(4, 90))
    fmt.Println(math.Min(4, 90))

    fmt.Println("--- 小数点以下の切り捨て切り上げ ---")
    fmt.Println(math.Trunc(1.5))
    fmt.Println(math.Trunc(-1.5))
    fmt.Println(math.Floor(3.14))
    fmt.Println(math.Floor(-3.14))
    fmt.Println(math.Ceil(3.14))
    fmt.Println(math.Ceil(-3.14))

    {
        fmt.Println("--- 非数と無限大 ---")
        n := math.Sqrt(-1)
        fmt.Printf("Sqrt(-1)=%v\n", n)
        fmt.Println(math.IsNaN(n))

        fmt.Println(math.Inf(0))
        fmt.Println(math.Inf(-1))
        fmt.Println(math.NaN())
    }
}
