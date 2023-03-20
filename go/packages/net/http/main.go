package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "net/url"
    "os"
)

func main() {
    {
        fmt.Println("--- GETメソッド ---")
        res, err := http.Get("https://www.google.com")
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(res.StatusCode)
        fmt.Println(res.Proto)
        fmt.Println(res.Header["Date"])
        fmt.Println(res.Header["Content-Type"])
        fmt.Println(res.Request.Method)
        fmt.Println(res.Request.URL)

        defer res.Body.Close()

        body, err := ioutil.ReadAll(res.Body)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(string(body))
    }
    {
        fmt.Println("--- POSTメソッドによるフォームの送信 ---")
        vs := url.Values{}
        vs.Add("id", "1")
        vs.Add("msg1", "message-text")
        vs.Add("msg2", "メッセージテキスト")
        fmt.Println(vs.Encode())

        /*res, err := http.PostForm("https://example.com/comments/post", vs)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(res.StatusCode)*/
    }
    {
        fmt.Println("--- ファイルのアップロード ---")
        f, err := os.Open("foo.jpeg")
        if err != nil {
            log.Fatal(err)
        }

        fi, err := f.Stat()
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(fi.Name())

        /*res, err := http.Post("https://example.com/upload", "image/jpeg", f)
        if err != nil {
            log.Fatal(err)
        }*/
    }
}
