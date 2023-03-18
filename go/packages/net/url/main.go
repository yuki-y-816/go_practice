package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	{
		fmt.Println("--- URLをパースする ---")

		u, err := url.Parse("https://www.example.com/search?name=yuki&age=33#top")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(u.Scheme)
		fmt.Println(u.Host)
		fmt.Println(u.Path)
		fmt.Println(u.RawQuery)
		fmt.Println(u.Fragment)

		fmt.Println(u.IsAbs())
		fmt.Println(u.Query())
	}
	{
		fmt.Println("--- URLを生成する ---")

		u := &url.URL{}

		u.Scheme = "https"
		u.Host = "google.com"

		q := u.Query()
		q.Set("query", "GO言語")

		u.RawQuery = q.Encode()

		fmt.Println(u)
	}
}
