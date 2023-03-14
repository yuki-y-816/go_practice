package main

import (
	"log"
	"os"
)

func main() {
	log.Print("ログの1行目\n")
	log.Println("ログの2行目")
	log.Printf("ログの%d行目\n", 3)

	f, err := os.OpenFile("test.log", os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	log.SetOutput(f)
	log.Println("ログのメッセージ")

	log.SetPrefix("[LOG] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("ログのメッセージ2")

	logger := log.New(os.Stdout, "[NEW LOG] ", log.LstdFlags)
	logger.Fatalln("ログの出力の停止")
}
