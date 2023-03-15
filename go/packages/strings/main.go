package main

import (
	"fmt"
	"strings"
)

func main() {
	{
		fmt.Println("--- strings.Join() ---")
		fmt.Println(strings.Join([]string{"Hello", ",", "World", "."}, " "))
		fmt.Println(strings.Join([]string{"A", "B", "C", "D"}, ","))
	}
	{
		fmt.Println("--- strings.Index() ---")
		str := "ABCDEFG"
		fmt.Println(strings.Index(str, "A"))
		fmt.Println(strings.Index(str, "E"))
		fmt.Println(strings.Index(str, "DEF"))
		fmt.Println(strings.Index(str, "Y"))

		fmt.Println("--- strings.LastIndex() ---")
		fmt.Println(strings.LastIndex("ABCABCABCABC", "ABC"))

		fmt.Println("--- strings.IndexAny() ---")
		fmt.Println(strings.IndexAny(str, "ABC"))
		fmt.Println(strings.IndexAny(str, "BDX"))
		fmt.Println(strings.IndexAny(str, "XYZ"))

		fmt.Println("--- strings.LastIndexAny() ---")
		fmt.Println(strings.LastIndexAny(str, "ACF"))
		fmt.Println(strings.LastIndexAny(str, "XYZ"))

		fmt.Println("--- strings.Contains() ---")
		fmt.Println(strings.Contains(str, "AB"))
		fmt.Println(strings.Contains(str, "CDF"))
		fmt.Println(strings.Contains(str, "X"))
		fmt.Println(strings.Contains(str, "ACE"))

		fmt.Println("--- strings.ContainsAny() ---")
		fmt.Println(strings.ContainsAny(str, "ABC"))
		fmt.Println(strings.ContainsAny(str, "XYZ"))
		fmt.Println(strings.ContainsAny(str, "AHR"))
	}
	{
		str := "GO言語"
		fmt.Println("--- strings.HasPrefix() ---")
		fmt.Println(strings.HasPrefix(str, "GO"))
		fmt.Println(strings.HasPrefix(str, "言語"))

		fmt.Println("--- strings.HasSuffix() ---")
		fmt.Println(strings.HasSuffix(str, "GO"))
		fmt.Println(strings.HasSuffix(str, "言語"))
	}
	{
		str := "AAAAAABBBBBBBBBCCCCCCCCDDD"
		fmt.Println("--- strings.Count() ---")
		fmt.Println(strings.Count(str, "A"))
		fmt.Println(strings.Count(str, "C"))
		fmt.Println(strings.Count("ABCABCABCABD", "ABC"))
	}
	{
		fmt.Println("--- strings.Repeat() ---")
		fmt.Println(strings.Repeat("ABC", 4))
		fmt.Println(strings.Repeat("ABC", 0))
	}
	{
		str := "AAAAAAA"
		fmt.Println("--- strings.Replace() ---")
		fmt.Println(strings.Replace(str, "A", "Hello", 2))
		fmt.Println(strings.Replace(str, "A", "Hello", -1))
		fmt.Println(strings.Replace("C言語", "C", "GO", 1))
	}
	{
		str := "A,B,C,D,E,F"
		fmt.Println("--- strings.Split() ---")
		fmt.Println(strings.Split(str, ","))
		fmt.Println("--- strings.SplitAfter() ---")
		fmt.Println(strings.SplitAfter(str, ","))
		fmt.Println("--- strings.SplitN() ---")
		fmt.Println(strings.SplitN(str, ",", 3))
		fmt.Println("--- strings.SplitAfterN() ---")
		fmt.Println(strings.SplitAfterN(str, ",", 3))
	}
	{
		fmt.Println("--- strings.ToLower() ---")
		fmt.Println(strings.ToLower("Hello, World!"))
		fmt.Println("--- strings.ToUpper() ---")
		fmt.Println(strings.ToUpper("Hello, World!"))
	}
	{
		fmt.Println("--- strings.TrimSpace() ---")
		fmt.Println(strings.TrimSpace(" - hello, world - "))
		fmt.Println(strings.TrimSpace("\t- hello, world -\r\n"))
		fmt.Println(strings.TrimSpace("　-日本語の全角-　"))

		fmt.Println("--- strings.Fields() ---")
		fmt.Println(strings.Fields("a b  c"))
		fmt.Println(strings.Fields(" A\t B\tC \nD\rE"))
		fmt.Println(strings.Fields("い　　ろ　は　　にほ　へ　　　　　と"))
	}
}
