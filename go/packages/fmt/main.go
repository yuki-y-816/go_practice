package main

import (
	"fmt"
)

type User struct {
	Id    int
	Email string
}

func main() {
	u := &User{Id: 7, Email: "sample@email.com"}

	fmt.Printf("%v\n", *u)
	fmt.Printf("%+v\n", *u)
	fmt.Printf("%#v\n", *u)

	fmt.Println(123, 456, "Golang", "is great lang!!", struct{ X, Y int }{100, 99})
}
