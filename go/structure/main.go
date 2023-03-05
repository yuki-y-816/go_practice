package main

import (
	f "fmt"

	"structure/definition"
	"structure/function"
	"structure/mytype"
	"structure/slice"
)

func main() {
	mytype.Main()
	definition.Main()
	function.Main()

	f.Println("-- my constructor! --")
	f.Println(function.NewUser(7, "Taro"))

	slice.Main()
}
