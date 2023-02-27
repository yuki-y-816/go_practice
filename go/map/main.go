package main

import (
	f "fmt"
)

var simpleMap map[int]string

func createMap() {
	f.Println("--- createMap() ---")
	m := make(map[int]string)

	m[1] = "US"
	m[81] = "Japan"
	m[86] = "China"

	f.Println(m)

	m1 := make(map[string]string)

	m1["Yamada"] = "Taro"
	m1["Sato"] = "Hanako"
	m1["Yamada"] = "Jiro"

	f.Println(m1)
}

func mapLiteral() {
	f.Println("--- mapLiteral() ---")

	m1 := map[int]string{1: "Taro", 2: "Hanako", 3: "Jiro"}
	f.Println(m1)

	m2 := map[int]string{
		1: "Taro",
		2: "Hanako",
		3: "Jiro", // カンマが必要
	}
	f.Println(m2)

	m3 := map[int][]int{
		1: []int{1},
		2: []int{1, 2},
		3: []int{1, 2, 3},
		4: {1, 2, 3, 4},
	}
	f.Println(m3)

	m4 := map[int]map[float64]string{
		1: {
			3.14: "円周率",
			5.00: "整数",
		},
	}
	f.Println(m4)
}

func referMap() {
	f.Println("--- referMap() ---")

	m := map[int]string{1: "A", 2: "B", 3: "C"}
	f.Printf("m[3]=%v\n", m[3])
	f.Printf("m[9]=%v\n", m[9]) // 存在しないキーから取り出すと初期値が入る（ string の初期値は "" ）

	{
		s, ok := m[1]
		f.Printf("m[1]=%v, ok=%v\n", s, ok)
	}
	{
		s, ok := m[5]
		f.Printf("m[5]=%v, ok=%v\n", s, ok)
	}

	if _, ok := m[3]; ok {
		f.Println("m[3]が存在するときの処理")
	}
	if _, ok := m[5]; ok {
		f.Println("m[5]が存在するときの処理")
	} else {
		f.Println("m[5]が存在しないときの処理")
	}
}

func mapRange() {
	f.Println("--- mapRange() ---")

	m := map[int]string{
		1: "Apple",
		2: "Banana",
		3: "Cherry",
	}
	for k, v := range m {
		f.Printf("key=%d, value=%s\n", k, v)
	}

	m1 := map[string]int{
		"Apple":  88,
		"Banana": 107,
		"Cherry": 46,
	}
	m1["Grape"] = 66
	m1["Lemon"] = 16
	m1["Orange"] = 44
	m1["Pineapple"] = 73
	for k, v := range m1 { // 順番はめちゃくちゃになる
		f.Printf("%s => %d\n", k, v)
	}
}

func mapLen() {
	f.Println("--- mapLen() ---")
	m := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}
	f.Printf("len(%v)=%d\n", m, len(m))
}

func deleteMap() {
	f.Println("--- deleteMap() ---")
	m := map[int]string{
		1: "A",
		2: "B",
		3: "C",
		4: "D",
	}
	f.Printf("m=%v\n", m)

	delete(m, 2)
	f.Printf("m=%v\n", m)
}

func main() {
	f.Printf("%T\n", simpleMap)

	createMap()
	mapLiteral()
	referMap()
	mapRange()
	mapLen()
	deleteMap()
}
