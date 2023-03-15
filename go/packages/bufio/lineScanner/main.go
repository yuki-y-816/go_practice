package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	{
		s := `XXXX
        YYYY
        ZZZZZZZZZZZZZZ`

		reader := strings.NewReader(s)
		scanner := bufio.NewScanner(reader)

		scanner.Scan()
		fmt.Println(scanner.Text())

		scanner.Scan()
		fmt.Println(scanner.Text())

		scanner.Scan()
		fmt.Println(scanner.Text())
	}
	{
		s := `ABC DEF HIJ
        GHI JKL MNO PQR
        STU VWX
        YZ`

		reader := strings.NewReader(s)
		scanner := bufio.NewScanner(reader)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}
