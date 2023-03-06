package main

import (
	f "fmt"
)

type MyError struct {
	Message string
	ErrCode int
}

func (e MyError) Error() string {
	return e.Message
}
func RaiseError() error {
	return &MyError{Message: "エラーが発生しました", ErrCode: 1234}
}

type Stringify interface {
	ToString() string
}
type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return f.Sprintf("%s(%d)", p.Name, p.Age)
}
func (p *Person) String() string {
	// "fmt"に定義されてる Stringer インターフェースに合わせてメソッドを定義すると
	// fmt.Println()に渡すだけでその内容で出力される
	return f.Sprintf("%s is %d years old!", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return f.Sprintf("[%s] %s", c.Number, c.Model)
}

func main() {
	err1 := RaiseError()
	f.Printf("%v\n", err1)
	f.Println(err1.Error())
	e1, ok1 := err1.(*MyError)
	if ok1 {
		f.Printf("err1.ErrCode=%v\n", e1.ErrCode)
	}

	vs := []Stringify{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "xxx-0123", Model: "PX512"},
	}
	for _, v := range vs {
		f.Println(v.ToString())
	}

	p := &Person{Name: "Hanako", Age: 19}
	f.Println(p)
}
