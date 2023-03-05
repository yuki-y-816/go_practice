package function

import (
	f "fmt"
	"math"
)

type Pointer struct {
	X, Y int
}

func structArgument() {
	f.Println("-- structArgument() --")
	p := Pointer{X: 77, Y: 88}

	showPointer(p)

	swapPointer(p)
	f.Println("after swapPointer()")
	showPointer(p)

	swapPointerPointer(&p)
	f.Println("after swapPointerPointer()")
	showPointer(p)
}
func showPointer(p Pointer) {
	f.Println(p)
}
func swapPointer(p Pointer) {
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}
func swapPointerPointer(p *Pointer) {
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}

func newPointer() {
	f.Println("--- newPointer() ---")
	p := new(Pointer)
	f.Printf("new(Pointer)=%v\n", p)

	p.X = 15
	p.Y = 22
	f.Println(*p)
}

// 原則として構造体のレシーバーはポインタ型で定義すべき
// メソッド内で構造体の変更をしてもその成果を得にくいため
func structFunc() {
	f.Println("--- structFunc() ---")

	p := Pointer{X: 110, Y: 6}

	p.render()
	dis := p.distance(&Pointer{X: 100, Y: 1})
	f.Printf("dis=%v\n", dis)
}
func (p *Pointer) render() {
	f.Printf("<%d, %d>\n", p.X, p.Y)
}
func (p *Pointer) distance(dp *Pointer) float64 {
	// レシーバーの型（(p *Pointer)の部分）が違ってれば同名のメソッドを別に定義することも可
	x, y := p.X-dp.X, p.Y-dp.Y

	return math.Sqrt(float64(x*x + y*y))
}

type MyInt int

func (m MyInt) plus(i int) int {
	return int(m) + i
}

type MyIntPair [2]int

func (m MyIntPair) first() int {
	return m[0]
}
func (m MyIntPair) last() int {
	return m[1]
}
func (m MyIntPair) toString() string {
	return f.Sprintf("[%d, %d]\n", m[0], m[1])
}
func aliasMethods() {
	f.Println("--- aliasMethod() ---")

	f.Printf("MyInt(4).plus(6)=%d\n", MyInt(4).plus(6))

	ip := MyIntPair{3, 8}
	f.Printf("ip.first=%d\n", ip.first())
	f.Printf("ip.last=%d\n", ip.last())

	ms := MyString{"Hello", "my", "name", "is", "Yuki"}
	f.Printf("string joined -> %s\n", ms.join(","))
}

type MyString []string

func (ms MyString) join(sep string) string {
	sum := ""
	for _, v := range ms {
		if sum != "" {
			sum += sep
		}

		sum += v
	}

	return sum
}

type User struct {
	Id   int
	Name string
}

func NewUser(id int, name string) *User {
	u := new(User)
	u.Id = id
	u.Name = name

	return u
}

func Main() {
	structArgument()
	newPointer()
	structFunc()
	aliasMethods()

	f.Println("--- MyIntPair.toString() ---")
	funcVar := MyIntPair{3, 4}.toString
	f.Println(funcVar())
}
