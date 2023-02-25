package main

import (
	"fmt"
	"strconv"
)

var arr [10]int

func assignAndRef() {
	fmt.Println("--- assignAndRef() ---")

	slice := make([]float64, 3)
	fmt.Printf("%v\n", slice)

	slice[0] = 3.14
	fmt.Printf("%v\n", slice)

	slice[1] = 6.28
	fmt.Printf("%v\n", slice)

	// fmt.Println(slice[4]) // 要素数を越してるのでエラー
}

func myLen() {
	fmt.Println("--- myLen() ---")
	slice := make([]int, 8)

	fmt.Println(len(slice))
	fmt.Println(len([4]int{11, 22, 33, 44}))
}

func myCap() {
	fmt.Println("--- myCap() ---")

	s1 := make([]int, 5)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	s2 := make([]int, 5, 10)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
}

func simpleSlice() {
	fmt.Println("--- simpleSlice() ---")
	slice := [...]int{10, 20, 30, 40, 50}

	fmt.Println(slice)
	fmt.Printf("slice[2:4] : %v\n", slice[2:4])
	fmt.Printf("slice[:3] : %v\n", slice[:3])
	fmt.Printf("slice[:] : %v\n", slice[:])
	fmt.Printf("slice[:len()-2] : %v\n", slice[:len(slice)-2])

	fmt.Printf("'ABCDEFG'[3:6] : %v\n", "ABCDEFG"[3:6])
	fmt.Printf("'あいうえお'[3:6] : %v\n", "あいうえお"[3:9])
}

func myAppend() {
	fmt.Println("--- myAppend() ---")
	s := []int{1, 2, 3}

	s = append(s, 4)
	fmt.Println(s)
	s = append(s, 5, 6)
	fmt.Println(s)
	s = append(s, []int{7, 8, 9, 10}...) // 別のスライスを渡す時は ... が必要
	fmt.Println(s)

	var b []byte
	b = append(b, "あいうえお"...)
	b = append(b, "かきくけこ"...)
	fmt.Println(b)

	s1 := make([]int, 0, 0)
	fmt.Printf("(A) len=%d, cap=%d\n", len(s1), cap(s1))
	s1 = append(s1, 1)
	fmt.Printf("(B) len=%d, cap=%d\n", len(s1), cap(s1))
	s1 = append(s1, []int{2, 3, 4}...)
	fmt.Printf("(C) len=%d, cap=%d\n", len(s1), cap(s1))
	s1 = append(s1, 5) // 容量が足りなくなると容量が増える
	fmt.Printf("(D) len=%d, cap=%d\n", len(s1), cap(s1))
	s1 = append(s1, 6, 7, 8, 9) // 容量が足りなくなると容量が増える
	fmt.Printf("(E) len=%d, cap=%d\n", len(s1), cap(s1))

	s2 := make([]int, 1024, 1024)
	fmt.Printf("len=%d, cap=%d\n", len(s2), cap(s2))
	s2 = append(s2, 0) // 容量は倍増するわけではなくてそのときのランタイムによる
	fmt.Printf("len=%d, cap=%d\n", len(s2), cap(s2))
}

func myCopy() {
	fmt.Println("--- myCopy() ---")

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{10, 11}
	n1 := copy(s1, s2)
	fmt.Printf("n1=%d, s1=%v\n", n1, s1)

	s3 := []int{6, 7, 8, 9, 10, 11, 12, 13}
	n2 := copy(s1, s3) // コピー先の要素数まで
	fmt.Printf("n2=%d, s1=%v\n", n2, s1)
}

func completeSlice() {
	fmt.Println("--- completeSlice() ---")
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)

	a1 := a[2:4:4]
	fmt.Printf("a[2:4:4]=%v, len=%d, cap=%d\n", a1, len(a1), cap(a1))
	a2 := a[2:4:6]
	fmt.Printf("a[2:4:6]=%v, len=%d, cap=%d\n", a2, len(a2), cap(a2))
}

func sliceAndFor() {
	fmt.Println("--- sliceAndFor() ---")
	s := []string{"Apple", "Banana", "Cherry"}

	for i, v := range s {
		fmt.Printf("[%d] = %s\n", i, v)
	}

	for i, v := range s {
		fmt.Printf("[%d] = %s\n", i, v)
		s = append(s, "Melon"+strconv.Itoa(i))
	}
	fmt.Println(s)
}

func sum(s ...int) int {
	fmt.Printf("--- sum(%v) ---\n", s)

	n := 0
	for _, v := range s {
		n += v
	}

	return n
}

func powArr(a [3]int) {
	for i, v := range a {
		a[i] = v * v
	}

	return
}

func powSlice(s []int) {
	for i, v := range s {
		s[i] = v * v
	}

	return
}

func main() {
	fmt.Printf("var [10]int : %T\n", arr)
	fmt.Printf("var [10]int = %v\n", arr)                 // 配列
	fmt.Printf("make([]int, 10) = %v\n", make([]int, 10)) // スライス

	newS := [4]int{11, 22, 33, 44}
	fmt.Printf("newS : %v, len() : %v, cap() : %v\n", newS, len(newS), cap(newS))

	assignAndRef()
	myLen()
	myCap()
	simpleSlice()
	myAppend()
	myCopy()
	completeSlice()
	sliceAndFor()

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(sum([]int{1, 2, 3, 4, 5}...)) // スライスに ... を付けて引数に渡すと可変長引数として渡せる

	fmt.Println("--- powArr() ---")
	arrr := [3]int{1, 2, 3}
	powArr(arrr)
	fmt.Println(arrr)
	fmt.Println("--- powSlice() ---")
	sli := []int{1, 2, 3}
	powSlice(sli) // スライスは参照型！
	fmt.Println(sli)
}
