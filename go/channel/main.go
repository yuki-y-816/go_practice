package main

import (
	f "fmt"
	"time"
)

func basicFunc() {
	f.Println("--- basicFunc() ---")
	ch := make(chan int, 8)

	ch <- 5
	ch <- 7
	i1 := <-ch
	f.Printf("i1=%d\n", i1)
	i2 := <-ch
	f.Printf("i2=%d\n", i2)
}

func receiver(ch <-chan int) {
	for {
		i := <-ch
		f.Println(i)
	}
}

func basicGoroutin() {
	f.Println("--- basicGoroutin() ---")
	ch := make(chan int)

	go receiver(ch)

	i := 1
	for i <= 15 {
		ch <- i
		i++
	}
}

func chanLen() {
	ch := make(chan string, 3)

	ch <- "Apple"
	ch <- "Banana"
	f.Printf("len=%d\n", len(ch))
	ch <- "Cherry"
	f.Printf("len=%d\n", len(ch))
}

func chanCap() {
	ch1 := make(chan int, 4)
	ch2 := make(chan int, 99)

	f.Printf("cap(ch1)=%d\n", cap(ch1))
	f.Printf("cap(ch2)=%d\n", cap(ch2))
}

func chanClose() {
	f.Println("--- chanClose() ---")
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)

	var (
		i  int
		ok bool
	)

	i, ok = <-ch
	f.Printf("i=%d, ok=%v\n", i, ok)
	i, ok = <-ch
	f.Printf("i=%d, ok=%v\n", i, ok)
	i, ok = <-ch
	f.Printf("i=%d, ok=%v\n", i, ok)
	i, ok = <-ch
	f.Printf("i=%d, ok=%v\n", i, ok) // close されてバッファ内に何もないと初期値と false が返る
}

func receiver2(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			break
		}

		f.Printf("%s -> %d\n", name, i)
	}
	f.Printf("%s is done\n", name)
}

func threeGoRoutine() {
	f.Println("--- threeGoRoutine() ---")
	ch := make(chan int, 20)

	go receiver2("1st goroutine", ch)
	go receiver2("2nd goroutine", ch)
	go receiver2("3rd goroutine", ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}

	close(ch)

	time.Sleep(3 * time.Second)
}

func sampleChannelGoroutin() {
	f.Println("--- sampleChannelGoroutin() ---")

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for {
			i := <-ch1
			ch2 <- (i * 2)
		}
	}()

	go func() {
		for {
			i := <-ch2
			ch3 <- (i - 1)
		}
	}()

	n := 1
LOOP:
	for {
		select {
		case ch1 <- n:
			n++
		case i := <-ch3:
			f.Printf("recieved %d\n", i)
		default:
			if n > 100 {
				break LOOP
			}
		}
	}
}

func main() {
	basicFunc()
	basicGoroutin()
	chanLen()
	chanCap()
	chanClose()
	threeGoRoutine()
	sampleChannelGoroutin()
}
