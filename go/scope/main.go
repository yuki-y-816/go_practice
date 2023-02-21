package main

import "fmt"

// 分割パッケージは go build してないと使えない

func appName() string {
    const AppName = "Go Application"
    var Version = "1.0"

    return AppName + " " + Version
}

func main() {
    fmt.Println("AppVersion:", AppVersion)
    printMessage("hello!!!")
    printPi()
    fmt.Println(appName())
    //fmt.Println(AppName)  // appName() 内のスコープなのでコンパイルエラー
}
