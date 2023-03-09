package main

import (
	"fmt"
	"os"
)

func allEnv() {
	fmt.Println("--- allEnv() ---")
	for _, v := range os.Environ() {
		fmt.Println(v)
	}
}

func chkHome() {
	if home, ok := os.LookupEnv("HOME"); ok {
		fmt.Println(home)
	} else {
		fmt.Println("no $HOME")
	}
}

func main() {
	{
		allEnv()

		fmt.Println(os.Getenv("GOLANG_VERSION"))

		os.Clearenv()
		allEnv()
		chkHome()

		os.Setenv("HOME", "/home/is/where")
		allEnv()
		chkHome()
	}
	{
		fmt.Println(os.Getpid())
		fmt.Println(os.Getppid())
		fmt.Println(os.Getuid())
		fmt.Println(os.Geteuid())
		fmt.Println(os.Getgid())
		fmt.Println(os.Getegid())
	}
}
