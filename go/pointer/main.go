package main

import (
	f "fmt"
)

func increment(p *int) { // ポインタを引数で渡すときは 演算子* つけてデータ本体を渡す
	*p++
}

func pow(p *[3]int) {
	i := 0

	for i < 3 {
		(*p)[i] = (*p)[i] * (*p)[i]
		i++
	}
}

func basicFunc() {
	f.Println("--- basicFunc() ---")
	// ただの宣言は nil が入る
	var i0 *int
	f.Printf("i=%v\n", i0)

	// 変数の前に 演算子& をつけることでポインタ型を生成
	var i int
	p := &i
	f.Printf("p=%T\n", p)
	pp := &p
	f.Printf("pp=%T\n", pp)

	// 演算子* を変数の前につけることでポインタの示すメモリ上のアドレス経由でデータの本体をデリファレンス
	i = 77
	f.Printf("p=%v\n", p)
	f.Printf("*p=%v\n", *p)
	*p = 1000
	f.Printf("i=%v\n", i)

	increment(p)
	f.Printf("incremented *p=%v\n", *p)
	f.Printf("incremented i=%v\n", i)
	increment(p)
	f.Printf("incremented *p=%v\n", *p)
	f.Printf("incremented i=%v\n", i)
	increment(&i)
	f.Printf("incremented *p=%v\n", *p)
	f.Printf("incremented i=%v\n", i)

	arr := &[3]int{1, 2, 3}
	pow(arr)
	f.Printf("powed *arr : %v\n", *arr)

	// 具体的なアドレス
	f.Printf("p is type=%T, address=%p\n", p, p)
}

func pointerToArray() {
	f.Println("--- pointerToArray() ---")

	// 配列のポインタは (*p)[0] と書かなくても自動でデリファレンスして展開してくれてる

	p := &[3]int{1, 2, 3}
	f.Printf("(*p)[1]=%d\n", (*p)[1])
	f.Printf("p[0]=%v\n", p[0])

	sa := [3]string{"Apple", "Banana", "Grape"}
	p1 := &sa
	f.Printf("sa[1] = %s\n", sa[1])
	f.Printf("p1[1] = %s\n", p1[1])
	p1[1] = "SuperGrape"
	f.Printf("sa[1] = %s\n", sa[1])
	f.Printf("p1[1] = %s\n", p1[1])

	f.Printf("len(p1)=%d\n", len(p1))
	f.Printf("cap(p1)=%d\n", cap(p1))
	f.Printf("p1[0:2]=%v\n", p1[0:2])

	for i, v := range p1 {
		f.Printf("for range p1[%d]=%s\n", i, v)
	}
}

func stringType() {
	f.Println("--- stringType() ---")

	s := "ABCD"
	f.Printf("s=%s, &s=%v\n", s, &s)
	f.Printf("s[0]=%v\n", s[0])

	a := []string{"E", "F", "G"}
	for _, v := range a {
		s = s + v
		f.Printf("s=%s, &s=%v\n", s, &s)
	}

	printString(s)
	f.Printf("s=%s, &s=%v\n", s, &s)

}

func printString(s string) {
	f.Println("--- printString() ---")
	f.Printf("passed s=%s, &s=%v\n", s, &s)
}

func main() {
	basicFunc()
	pointerToArray()
	stringType()
}
