package main

import (
    "fmt"
    "os"
    "log"
)

func main() {
    f, err := os.Open("foo.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    bs := make([]byte, 128)
    {
        n, err := f.Read(bs)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(string(bs[:n]))
    }

    {
        fi, err := f.Stat()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("name=%v\n",fi.Name())
        fmt.Printf("size=%v\n",fi.Size())
        fmt.Printf("mode=%v\n",fi.Mode())
        fmt.Printf("mod_time=%v\n",fi.ModTime())
        fmt.Printf("is_dir=%v\n",fi.IsDir())
        fmt.Println("-----------------")
    }

    {
        new, _ := os.Create("bar.txt")

        fi, _ := new.Stat()
        bs := make([]byte, 128)
        fmt.Printf("name=%v\n", fi.Name())
        fmt.Printf("size=%v\n", fi.Size())
        fmt.Printf("is_dir=%v\n", fi.IsDir())

        new.Write([]byte("Hello, world!\n"))
        new.WriteAt([]byte("Golang"), 7)
        new.Seek(0, os.SEEK_END)
        new.WriteString("Yeah")

        f, err1 := os.Open("bar.txt")
        if err1 != nil {
            log.Fatal(err1)
        }

        n, _ := f.Read(bs)
        fmt.Println(string(bs[:n]))
        fmt.Println("-----------------")
        f.Close()

        os.Mkdir("bar", 0777)
        os.Rename("bar.txt", "bar/bar.txt")
        moved, err2 := os.Open("bar/bar.txt")
        if err2 != nil {
            log.Fatal(err2)
        }
        fi2, _ := moved.Stat()
        fmt.Printf("name=%v\n", fi2.Name())
        fmt.Printf("size=%v\n", fi2.Size())
        fmt.Printf("is_dir=%v\n", fi2.IsDir())
        fmt.Println("-----------------")
        moved.Close()

        err := os.RemoveAll("bar")
        if err != nil {
            log.Fatal(err)
        }
    }
}
