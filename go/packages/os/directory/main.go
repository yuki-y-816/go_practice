package main

import (
	"fmt"
	"log"
	"os"
)

func raiseError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	{
		dir, err := os.Getwd()
		raiseError(err)

		fmt.Println(dir)
	}
	{
		err1 := os.Chdir("../")
		raiseError(err1)
		dir, err2 := os.Getwd()
		raiseError(err2)

		fmt.Println(dir)
		err3 := os.Chdir("./directory")
		raiseError(err3)
	}
	{
		f, err1 := os.Open(".")
		raiseError(err1)

		defer f.Close()

		fis, err2 := f.Readdir(0)
		raiseError(err2)

		for _, fi := range fis {
			if fi.IsDir() {
				fmt.Printf("%s is a dirctory\n", fi.Name())

			} else {
				fmt.Printf("%s is a file\n", fi.Name())
			}
		}
	}
	{
		err2 := os.MkdirAll("foo/bar/baz", 0775)
		raiseError(err2)
	}
	{
		fmt.Println(os.TempDir())
	}
	{
		_, err0 := os.Stat("./foo/bar/baz/baz.txt")
		if !os.IsNotExist(err0) {
			err1 := os.Symlink("./foo.txt", "./foo/bar/baz/baz.txt")
			raiseError(err1)
		}

		paths, err2 := os.Readlink("./foo/bar/baz/baz.txt")
		fmt.Println(paths)
		raiseError(err2)
	}
	{
		host, err := os.Hostname()
		raiseError(err)
		fmt.Println(host)
	}
}
