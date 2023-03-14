package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("X is digit -> %t\n", unicode.IsDigit('X'))
	fmt.Printf("3 is digit -> %t\n", unicode.IsDigit('3'))
	fmt.Printf("３(zenkaku) is digit -> %t\n", unicode.IsDigit('３'))
	fmt.Printf("三 is digit -> %t\n", unicode.IsDigit('三'))

	fmt.Printf("A is letter -> %t\n", unicode.IsLetter('A'))
	fmt.Printf("3 is letter -> %t\n", unicode.IsLetter('3'))
	fmt.Printf("あ is letter -> %t\n", unicode.IsLetter('あ'))
	fmt.Printf("& is letter -> %t\n", unicode.IsLetter('&'))

	fmt.Printf("' ' is space -> %t\n", unicode.IsSpace(' '))
	fmt.Printf("\\t is space -> %t\n", unicode.IsSpace('\t'))
	fmt.Printf("_ is space -> %t\n", unicode.IsSpace('_'))

	fmt.Printf("\\n is control -> %t\n", unicode.IsControl('\n'))
	fmt.Printf("\\t is control -> %t\n", unicode.IsControl('\t'))
}
