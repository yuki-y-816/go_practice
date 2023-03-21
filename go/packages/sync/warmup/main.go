package main

import (
	"fmt"
	"time"
)

var myStruct struct{ A, B, C int }

func UpdateAndPrint(n int) {
	myStruct.A = n
	time.Sleep(time.Microsecond)

	myStruct.B = n
	time.Sleep(time.Microsecond)

	myStruct.C = n
	time.Sleep(time.Microsecond)

	fmt.Println(myStruct.A, myStruct.B, myStruct.C)
}

func main() {
	for i := 0; i < 300; i++ {
		go func() {
			for i := 0; i < 300; i++ {
				UpdateAndPrint(i)
			}
		}()
	}

	for {
	}
}
