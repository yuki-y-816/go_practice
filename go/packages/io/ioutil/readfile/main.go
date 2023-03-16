package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	bs, err := ioutil.ReadFile("new-ocean-treaty.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(bs))
}
