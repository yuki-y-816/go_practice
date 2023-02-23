package constructs

import f "fmt"

func ForFunc() {
    i := 0
    for {
        f.Println(i)
        i++
        if i == 3 {
            break
        }
    }

    i = 0
    for i < 5 {
        f.Println(i)
        i++
    }

    for i := 0; i <= 10; i++ {
        f.Println("count to 10, now", i)
        i++
    }

    i = 0
    for ;i <= 5; {
        f.Println("count to 5, now", i)
        i++
    }

    for i := 0; i <= 10; i++ {
        if i % 2 == 0 {
            continue // 次のループ
        }

        if i == 9 {
            break // ループを抜ける
        }

        f.Println(i, "is odd")
    }

    fruits := [3]string {"Apple", "Banana", "Cherry"}
    for i, s := range fruits {
        f.Printf("fruits[%d] = %v\n", i, s)
    }
    for i, rune := range "ABCD" {
        f.Printf("[%d] : %v\n", i, rune)
    }
    for i, rune := range "あいうけこ" {
        f.Printf("[%d] : %v\n", i, rune)
    }

}
