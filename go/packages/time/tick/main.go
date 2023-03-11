package main

import (
	"fmt"
	"time"
)

func main() {
	ch := time.Tick(2 * time.Second)
	for {
		t := <-ch
		fmt.Println(t)
	}
}
