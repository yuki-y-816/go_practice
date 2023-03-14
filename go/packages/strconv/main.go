package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("--- bool -> string ---")
	fmt.Printf("bool -> %s\n", strconv.FormatBool(true))

	fmt.Println("--- int -> string ---")
	fmt.Printf("-12345 -> %s\n", strconv.FormatInt(-12345, 10))
	fmt.Printf("12345(2) -> %s\n", strconv.FormatInt(12345, 2))
	fmt.Printf("1179(16) -> %s\n", strconv.FormatInt(1179, 16))
	fmt.Printf("12345 -> %s\n", strconv.FormatUint(12345, 10))
	fmt.Printf("12345 -> %s\n", strconv.Itoa(12345))

	fmt.Println("--- float -> string ---")
	fmt.Printf("123.456 -> %s\n", strconv.FormatFloat(123.456, 'E', -1, 64))
	fmt.Printf("123.456 -> %s\n", strconv.FormatFloat(123.456, 'e', 2, 64))
	fmt.Printf("123.456 -> %s\n", strconv.FormatFloat(123.456, 'f', -1, 64))
	fmt.Printf("123.456 -> %s\n", strconv.FormatFloat(123.456, 'f', 2, 64))
	fmt.Printf("123.456 -> %s\n", strconv.FormatFloat(123.456, 'g', 4, 64))

	fmt.Println("--- string -> bool ---")
	bools := []string{
		"true",
		"1",
		"t",
		"TRUE",
		"false",
		"hello",
		"0",
		"f",
	}
	for _, v := range bools {
		b, err := strconv.ParseBool(v)

		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%s -> %t\n", v, b)
	}

	fmt.Println("--- string -> int ---")
	fmt.Println(strconv.ParseInt("1234", 10, 0))
	fmt.Println(strconv.Atoi("1234567"))
	fmt.Println(strconv.ParseInt("123456789", 10, 16))
	fmt.Println(strconv.ParseInt("1234", 8, 0))
	fmt.Println(strconv.ParseInt("-1234", 10, 0))
	fmt.Println(strconv.ParseUint("-1234", 10, 0))
	fmt.Println(strconv.ParseInt("7F4D", 16, 0))
	fmt.Println(strconv.ParseInt("123", 0, 0))
	fmt.Println(strconv.ParseInt("0123", 0, 0))
	fmt.Println(strconv.ParseInt("0x123", 0, 0))

	fmt.Println("--- string -> float ---")
	fmt.Println(strconv.ParseFloat("123.456", 64))
	fmt.Println(strconv.ParseFloat("-1.23456", 64))
	fmt.Println(strconv.ParseFloat(".456", 64))
	fmt.Println(strconv.ParseFloat("1.2345e8", 64))
	fmt.Println(strconv.ParseFloat("1.000000000001", 32))
	fmt.Println(strconv.ParseFloat("1.000000000001", 64))
	fmt.Println(strconv.ParseFloat("1e500", 64))
}
