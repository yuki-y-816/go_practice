package constructs

import (
    f "fmt"
    "runtime"
)

func subFunc(s string) {
    for i := 0; i < 5; i++ {
        f.Println(s);
    }
}

func GoFunc() {
    go subFunc("sub println") // 無限ループにしてると並行してるのがよくわかる

    f.Printf("NumCPU: %d\n", runtime.NumCPU())
    f.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
    f.Printf("Version: %s\n", runtime.Version())

    for i := 0; i < 3; i++ {
        f.Println("main go")
    }
}
