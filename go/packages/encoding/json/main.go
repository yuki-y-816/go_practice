package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type User struct {
	Id      int
	Name    string
	Email   string
	Created time.Time
}

func main() {
	{
		u := User{
			Id:      1,
			Name:    "山田太郎",
			Email:   "yamada@example.com",
			Created: time.Now(),
		}

		bs, err := json.Marshal(u)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(bs))
	}
	{
		src := `
            {
                "Id":12,
                "Name":"田中花子",
                "Email":"tanaka@example.com",
                "Created":"2023-02-18T11:34:00+09:00"
            }
        `

		u := new(User)

		err := json.Unmarshal([]byte(src), u)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", u)
	}
}
