package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Printf("now : %s\n", time.Now())

	{
		fmt.Println("-----------")
		t := time.Date(2023, 3, 10, 7, 28, 0, 0, time.Local)
		fmt.Println(t)
		fmt.Println(t.Year())
		fmt.Println(t.YearDay())
		fmt.Println(t.Month())
		fmt.Println(t.Weekday())
		fmt.Println(t.Day())
		fmt.Println(t.Hour())
		fmt.Println(t.Minute())
		fmt.Println(t.Second())
		fmt.Println(t.Nanosecond())
		fmt.Println(t.Zone())
	}
	{
		fmt.Println("-----------")
		fmt.Println(time.Hour)
		fmt.Println(time.Minute)
		fmt.Println(time.Second)
		fmt.Println(time.Millisecond)
		fmt.Println(time.Microsecond)
		fmt.Println(time.Nanosecond)

		d, _ := time.ParseDuration("3h30m")
		fmt.Println(d)

		t := time.Now()
		fmt.Println(t)
		fmt.Println(t.Add(3*time.Hour + 40*time.Minute))
	}
	{
		fmt.Println("--- 2020年オリンピックとの差 ---")
		dst := time.Date(2020, 7, 24, 0, 0, 0, 0, time.Local)
		now := time.Now()
		fmt.Println(dst.Sub(now))
	}
	{
		fmt.Println("--- 時刻の比較 ---")
		t0 := time.Date(2015, 10, 1, 0, 0, 0, 0, time.Local)
		t1 := time.Date(2015, 11, 1, 0, 0, 0, 0, time.Local)
		t2 := time.Date(2015, 11, 1, 0, 0, 0, 0, time.Local)

		fmt.Println(t0.Before(t1))
		fmt.Println(t0.After(t1))
		fmt.Println(t1.Equal(t2))
	}
	{
		fmt.Println("--- 年月日の増減 ---")
		now := time.Now()

		fmt.Println(now)
		fmt.Println(now.AddDate(1, 0, 0))
		fmt.Println(now.AddDate(0, -1, 0))
		fmt.Println(now.AddDate(0, -1, 0).AddDate(1, 1, 0))
	}
	{
		fmt.Println("--- 文字列から時刻を生成 ---")
		// 第1引数のフォーマットに、第2引数を完全に合わせないとダメ
		t1, err1 := time.Parse("2006/01/02", "1989/08/16") // 2006/01/02 じゃないとダメ
		if err1 != nil {
			log.Fatal(err1)
		}
		fmt.Println(t1)

		t2, _ := time.Parse(time.RFC822, "11 Mar 23 09:44 JST")
		fmt.Println(t2)
	}
	{
		fmt.Println("--- 時刻から文字列を生成 ---")
		t := time.Now()
		fmt.Println(t.Format(time.RFC822))
		fmt.Println(t.Format("2006年01月02日 15時04分05秒"))
		fmt.Println(t.Format("2006/01/02"))
		fmt.Println(t.UTC())
	}
	{
		fmt.Println("--- ローカルタイムへの変換 ---")
		t := time.Date(2022, 11, 19, 16, 22, 55, 0, time.UTC)
		fmt.Println(t)
		fmt.Println(t.Local())
	}
	{
		fmt.Println("--- UNIX時間 ---")
		t := time.Now()

		fmt.Println(t.Unix())

		unix := time.Unix(1679526842, 0)
		fmt.Println(unix)
	}
	{
		fmt.Println("--- sleep ---")
		time.Sleep(3 * time.Second)
		fmt.Println("awake!!!")
	}
	{
		fmt.Println("--- After ---")
		ch := time.After(3 * time.Second)
		fmt.Println(<-ch)
	}
}
