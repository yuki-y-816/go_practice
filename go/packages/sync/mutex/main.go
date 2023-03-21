package main

import (
	"fmt"
	"sync"
	"time"
)

var myStruct struct{ A, B, C int }
var mutex *sync.Mutex

func UpdateAndPrintln(n int) {
	mutex.Lock()

	myStruct.A = n
	time.Sleep(time.Microsecond)
	myStruct.B = n
	time.Sleep(time.Microsecond)
	myStruct.C = n
	time.Sleep(time.Microsecond)

	fmt.Println(myStruct.A, myStruct.B, myStruct.C)

	mutex.Unlock()
}

func main() {
	mutex = new(sync.Mutex)

	for i := 0; i < 200; i++ {
		go func() {
			for i := 0; i < 200; i++ {
				UpdateAndPrintln(i)
			}
		}()
	}
	for {
	}
}
