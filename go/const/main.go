package main

import "fmt"

const ONE = 1

const (
    X = "a string"
    Y
    Z = "other string"
)
const XY = X + " & " + Y
const (
    B1 = true
    B2 = 2 > 5
)
const (
    U64 = uint64(1234)
    OVERU64 = 18446744073709551615 + 1 // uint64 の限界を超えたもの
)

func one() (int, int) {
    const TWO = int(2)
    return ONE, TWO
}

const C = 4.7 + 1.3i

const (
    I1 = iota
    I2
    I3 = iota + 1
    I4 = "this is I4"
    I5 = iota
)

func main() {
    one, two := one()
    fmt.Printf("one = %v, two = %v\n", one, two)

    fmt.Printf("X = %v, Y = %v, Z = %v\n", X, Y, Z)
    fmt.Printf("XY = %v\n", XY)

    fmt.Printf("B1 = %v, B2 = %v\n", B1, B2)

    u64 := U64
    fmt.Printf("u64 : %T\n", u64)
    fmt.Printf("OVERU64B - 1 = %v\n",uint64(OVERU64 - 1))

    fmt.Printf("C = %v\n", C)

    fmt.Printf("I1 = %v, I2 = %v, I3 = %v, I4 = %v, I5 = %v\n", I1, I2, I3, I4, I5)
}
