package main

import (
	"fmt"
	"regexp"
)

func main() {
	{
		fmt.Println("--- 正規表現のフラグ ---")
		re := regexp.MustCompile(`(?i)aBc`)

		fmt.Println(re.MatchString("ABC"))
	}
	{
		fmt.Println("--- 幅を持たない正規表現のパターン ---")
		re := regexp.MustCompile(`(?i)^abc$`)

		fmt.Println(re.MatchString("ABC"))
		fmt.Println(re.MatchString("  ABC"))
	}
	{
		fmt.Println("--- 基本的な正規表現のパターン ---")
		var str string

		re1 := regexp.MustCompile(`bc`)

		str = "abcdefg"
		fmt.Println(str, re1.MatchString(str))
		str = "ABCDEFG"
		fmt.Println(str, re1.MatchString(str))

		re2 := regexp.MustCompile(`.`)

		str = "ABC"
		fmt.Println(str, re2.MatchString(str))
		str = "日本語"
		fmt.Println(str, re2.MatchString(str))
		str = "\n"
		fmt.Println(str, re2.MatchString(str))

		re3 := regexp.MustCompile(`abc|xyz`)
		str = "abc"
		fmt.Println(str, re3.MatchString(str))
		str = "xyz"
		fmt.Println(str, re3.MatchString(str))
	}
	{
		fmt.Println("--- 繰り返しを表す正規表現のパターン ---")
		var str string

		re1 := regexp.MustCompile(`a+b*`)
		str = "ab"
		fmt.Println(str, re1.MatchString(str))
		str = "a"
		fmt.Println(str, re1.MatchString(str))
		str = "aaaaaaaaabbbccc"
		fmt.Println(str, re1.MatchString(str))
		str = "bbb"
		fmt.Println(str, re1.MatchString(str))

		re2 := regexp.MustCompile(`A+?A+?X`)
		str = "AAX"
		fmt.Println(str, re2.MatchString(str))
		str = "AX"
		fmt.Println(str, re2.MatchString(str))
		str = "AAAAAAAAX"
		fmt.Println(str, re2.MatchString(str))
		str = "ABAX"
		fmt.Println(str, re2.MatchString(str))
	}
	{
		fmt.Println("--- 正規表現の文字クラス ---")
		var str string

		re1 := regexp.MustCompile(`(?i)[XYZ]`)
		str = "X"
		fmt.Println(str, re1.MatchString(str))
		str = "y"
		fmt.Println(str, re1.MatchString(str))
		str = "A"
		fmt.Println(str, re1.MatchString(str))

		re2 := regexp.MustCompile(`^[a-zA-Z0-9_]{3}$`)
		str = "AbC"
		fmt.Println(str, re2.MatchString(str))
		str = "007"
		fmt.Println(str, re2.MatchString(str))
		str = "Hello"
		fmt.Println(str, re2.MatchString(str))
		str = "A_z"
		fmt.Println(str, re2.MatchString(str))

		re3 := regexp.MustCompile(`[^a-zA-Z0-9]`)
		str = "ABC"
		fmt.Println(str, re3.MatchString(str))
		str = "007"
		fmt.Println(str, re3.MatchString(str))
		str = "日本語"
		fmt.Println(str, re3.MatchString(str))
	}
	{
		fmt.Println("--- Perl 由来の定義済みクラス ---")
		var str string

		re1 := regexp.MustCompile(`\d+`)
		str = "12345"
		fmt.Println(str, re1.MatchString(str))
		str = "X=7"
		fmt.Println(str, re1.MatchString(str))
		str = "XYZ"
		fmt.Println(str, re1.MatchString(str))

		re2 := regexp.MustCompile(`\w+`)
		str = "AbC"
		fmt.Println(str, re2.MatchString(str))
		str = "007"
		fmt.Println(str, re2.MatchString(str))
		str = "Hello"
		fmt.Println(str, re2.MatchString(str))
		str = "A_z"
		fmt.Println(str, re2.MatchString(str))
	}
	{
		fmt.Println("--- Unicode に対応した文字クラス ---")
		var str string

		re := regexp.MustCompile(`^\p{Katakana}+$`)
		str = "アイウエオ"
		fmt.Println(str, re.MatchString(str))
		str = "ｱｲｳｴｵ"
		fmt.Println(str, re.MatchString(str))
		str = "あいうえお"
		fmt.Println(str, re.MatchString(str))
	}
	{
		fmt.Println("--- 正規表現のグループ ---")
		var str string

		re := regexp.MustCompile(`(佐藤|鈴木)(太郎|花子)`)
		str = "佐藤太郎"
		fmt.Println(str, re.MatchString(str))
		str = "鈴木花子"
		fmt.Println(str, re.MatchString(str))
		str = "佐藤健"
		fmt.Println(str, re.MatchString(str))
	}
	{
		fmt.Println("--- 正規表現にマッチした文字列の取得 ---")
		var str string

		re := regexp.MustCompile(`\w+`)
		str = "abc xyz 999"

		fmt.Println(re.FindString(str))
		fmt.Println(re.FindAllString(str, 2))
		fmt.Println(re.FindAllString(str, -1))
	}
	{
		fmt.Println("--- 正規表現による文字列の分割 ---")
		re1 := regexp.MustCompile(`\s+`)
		str := "A B C   D\tE,F"

		fmt.Println(re1.Split(str, 3))
		fmt.Println(re1.Split(str, -1))
	}
	{
		fmt.Println("--- 正規表現による文字列の置換 ---")

		re1 := regexp.MustCompile(`佐藤`)
		fmt.Println(re1.ReplaceAllString("佐藤さんと田中さん", "鈴木"))
		fmt.Println(re1.ReplaceAllString("田中さん", "山下"))

		re2 := regexp.MustCompile(`、|。`)
		fmt.Println(re2.ReplaceAllString("私は、Go言語を使用する、プログラマーです。", ""))
	}
	{
		fmt.Println("--- 正規表現のグループによるサブマッチ ---")
		re1 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
		str := `
            00-1111-2222
            3333-44-5555-6
            777-777-777
            5-5-5
            777
        `
		ms := re1.FindAllStringSubmatch(str, -1)

		for _, v := range ms {
			fmt.Println(v)
		}
	}
	{
		fmt.Println("--- 正規表現のグループと置換処理 ---")
		re := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)

		fmt.Println(re.ReplaceAllString("TEL: 000-111-222", "$3--$2 $1"))
	}
}
