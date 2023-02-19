package main

import (
    "fmt"
    "math"
)

func main () {
    fmt.Printf("uint32 max value = %d\n", math.MaxUint32)
    fmt.Printf("float32 max value = %v\n", math.MaxFloat32)
    fmt.Printf("float64 max value = %v\n", math.MaxFloat64)

    i := 1
    fmt.Printf("i is %T\n", i)

    f64 := 1.0
    fmt.Printf("f64 = 1.0 is %T\n", f64)

    f32 := float32(1.0)
    fmt.Printf("f32 = float32(1.0) is %T\n", f32)

    zero := 0.0
    fmt.Printf("zero = 0.0 is %v\n", zero)

    pinf := 1.0 / zero
    fmt.Printf("pinf is %v\n", pinf)

    ninf := -1.0 / zero
    fmt.Printf("ninf is %v\n", ninf)

    nan := zero / zero
    fmt.Printf("zero / zero = %v\n", nan)

    c := complex(1.3, 4.2)
    fmt.Printf("c's real is %v\n", real(c))
    fmt.Printf("c's imag is %v\n", imag(c))

    s := "GOの文字列"
    fmt.Printf("%v\n", s)

    raw := `
        GO の
        RAW文字列リテラルによる
        複数行に渡る
        文字列\n
    `
    fmt.Printf("%v\n", raw)

    a := [5]int{0,1,2,3,4}
    fmt.Printf("%v\n", a)
    fmt.Printf("a[1] = %v\n", a[1])

    a1 := [...]string{"a", "b", "c"}
    fmt.Printf("%v\n", a1)
    a2 := [3]string{"d", "e", "f"}
    a1 = a2
    fmt.Printf("%v\n", a1)
    a1[0] = "A"
    a1[2] = "C"
    fmt.Printf("a1 = %v\n", a1)
    fmt.Printf("a2 = %v\n", a2)

    var x interface{}
    fmt.Printf("interface{} is %v\n", x)
    x = 1
    fmt.Printf("interface{} = 1 is %v\n", x)
    x = "text"
    fmt.Printf("interface{} = 'text' is %v\n", x)
    x = [...]uint64{2, 100, 777}
    fmt.Printf("interface{} = [] is %v\n", x)
}
