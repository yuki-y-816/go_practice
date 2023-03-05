package mytype

import (
	f "fmt"
)

func basic() {
	f.Println("--- mytype.basic() ---")

	type MyInt int

	var n1 MyInt = 5
	n2 := MyInt(77)

	f.Printf("n1=%d\n", n1)
	f.Printf("n2=%d\n", n2)
}

func variousAlias() {
	f.Println("--- variousAlias() ---")
	type (
		IntPair     [2]int
		Strings     []string
		AreaMap     map[string][2]float64
		IntsChannel chan []int
	)

	f.Printf("IntPair : %v\n", IntPair{1, 2})
	f.Printf("Strings : %v\n", Strings{"A", "B", "C", "D"})
	f.Printf("AreaMap : %v\n", AreaMap{"Tokyo": {35.689488, 139.691706}})
	f.Printf("IntsChannel : %v\n", make(IntsChannel))
}

type ReturnInt func(i int) int

func Sum(ints []int, callback ReturnInt) int {
	var sum int
	for _, v := range ints {
		sum += v
	}

	return callback(sum)
}

func Main() {
	basic()
	variousAlias()

	f.Println("--- myCallbackSum() ---")
	num := Sum(
		[]int{1, 2, 3, 4, 5},
		func(i int) int {
			return i * 2
		},
	)
	f.Printf("num=%d\n", num)
}
