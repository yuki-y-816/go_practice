package definition

import (
	f "fmt"
)

func basic() {
	f.Println("--- definition.basic() ---")
	type Point struct {
		X    int
		Y, Z int
	}

	var pt Point
	f.Printf("pt.X=%d\n", pt.X)
	f.Printf("pt.Y=%d\n", pt.Y)
	f.Printf("pt.Z=%d\n", pt.Z)

	pt.X = 2
	pt.Z = 99
	f.Printf("pt.X=%d\n", pt.X)
	f.Printf("pt.Y=%d\n", pt.Y)
	f.Printf("pt.Z=%d\n", pt.Z)

	pt1 := Point{
		X: 5,
		Y: 9,
	}
	f.Printf("pt1.X=%d\n", pt1.X)
	f.Printf("pt1.Y=%d\n", pt1.Y)
	f.Printf("pt1.Z=%d\n", pt1.Z)
}

func structureInStructure() {
	f.Println("--- structureInStructure() ---")
	type Feed struct {
		Name   string
		Amount uint
	}

	type Animal struct {
		Name string
		Feed Feed
	}

	monkey := Animal{
		Name: "monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}

	f.Printf("name=%s\n", monkey.Name)
	f.Printf("Feed.name=%s\n", monkey.Feed.Name)
	f.Printf("Feed.amount=%d\n", monkey.Feed.Amount)

	type Base struct {
		Id   int
		Name string
	}

	type A struct {
		Base // 共通のフィールド
		Name string
		Area string
	}

	type B struct {
		Base   // 共通のフィールド
		Title  string
		Bodies []string
	}

	a := A{
		Base: Base{Id: 17, Name: "Taro"},
		Name: "Taro",
		Area: "Tokyo",
	}

	b := B{
		Base:   Base{Id: 81, Name: "Hanako"},
		Title:  "no title",
		Bodies: []string{"A", "B"},
	}

	f.Printf("a.Id=%v\n", a.Id)
	f.Printf("a.Base.Name=%v\n", a.Base.Name)
	f.Printf("b.Title=%s\n", b.Title)
	f.Printf("b.Bodies[1]=%s\n", b.Bodies[1])
}

func defineMap() {
	f.Println("--- defineMap() ---")
	type User struct {
		Id   int
		Name string
	}

	m1 := map[int]User{
		1: {Id: 1, Name: "Taro"},
		2: {Id: 2, Name: "Hanako"},
	}
	for i, v := range m1 {
		f.Printf("mi[%d]=%v\n", i, v)
	}

	ms := map[int][]string{
		1: {"A", "B"},
		2: {"A", "B", "C"},
	}
	f.Printf("ms=%v\n", ms)

	mm := map[int]map[int]string{
		1: {
			0: "Apple",
			1: "Banana",
		},
		2: {0: "hello", 1: "World"},
	}
	f.Printf("mm=%v\n", mm)
}

func Main() {
	basic()
	structureInStructure()
	defineMap()
}
