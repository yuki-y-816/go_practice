package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())

	fmt.Println("--- シード値固定により毎回同じ結果 ---")
	rand.Seed(256)
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())

	fmt.Println("--- 現在時刻による疑似乱数 ---")
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100)) // 0 以上 100 未満で乱数生成
	fmt.Println(rand.Int())
	fmt.Println(rand.Int63())

	fmt.Println("--- 擬似乱数生成器の生成 ---")
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	fmt.Println(rnd.Intn(100))
	fmt.Println(rnd.Int63())
}
