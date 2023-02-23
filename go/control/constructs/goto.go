package constructs

import f "fmt"

func labelControlStructure() {
    Loop:
    for {
        for {
            for {
                f.Println("for の深いループ")
                break Loop
            }
            f.Println("A")
        }
        f.Println("B")
    }

    Loop2:
    for i := 1 ; i <= 3; i++ {
        f.Printf("start with i = %v\n", i)
        for j := 1; j <= 3; j++ {
            if j > 1 {
                continue Loop2
            }
            f.Printf("%v * %v = %v\n", i, j, i*j)
        }
        f.Println("実行されない")
    }
}

func deferFunc() {
    defer f.Println("defer function!!")
    defer f.Println("defer function??")
    defer func(){
        f.Println("1")
        f.Println("2")
        f.Println("3")
    }() // <-関数呼び出しの形式
}

func GotoFunc() {
    // !! 基本 goto文 は必要ない !!
    f.Println("start GotoFunc()")
    goto L // 関数内でのみジャンプできる
    f.Println("A")
    //some := "something" // 変数定義もまたげない
    L:
    f.Println("B")

    labelControlStructure()

    deferFunc()
}
