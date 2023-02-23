package constructs

import "fmt"

// プログラマーが独自に init() を定義することで main() の前に実行される
func init() {
    fmt.Println("done init()")
}

func init() {
    // init() は何回でも定義できるけど悪趣味なのでダメ
    fmt.Println("done init() twice")
}

func InitFunc() {
    fmt.Println("do InitFunc()")
}
