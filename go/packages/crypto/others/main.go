package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

func main() {
	{
		fmt.Println("--- SHA1 ---")
		s := sha1.New()
		io.WriteString(s, "ABCDE")
		fmt.Printf("%x\n", s.Sum(nil))
	}
	{
		fmt.Println("--- SHA256 ---")
		s := sha256.New()
		io.WriteString(s, "ABCDE")
		fmt.Printf("%x\n", s.Sum(nil))
	}
	{
		fmt.Println("--- SHA512 ---")
		s := sha512.New()
		io.WriteString(s, "ABCDE")
		fmt.Printf("%x\n", s.Sum(nil))
	}
}
