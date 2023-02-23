package constructs

import f "fmt"

func IfFunc() {
	x := 1
	y := 5

	if x == 1 {
		f.Println("x == 1!")
	}

	if x > 2 {
		f.Println("x > 2!")
	} else if y < 5 {
		f.Println("y < 5!")
	} else {
		f.Println("else!")
	}

	// if 1 {
	//     f.Println("non-boolean condition in if statement")
	// }
	if true {
		f.Println("boolean condition")
	}

	if x, y := 2, 6; x < y {
		// ※ if{}内外でスコープが違うため x, y が上の定義と違ってる
		f.Printf("x(%d) is less than y(%d)\n", x, y)
	}

	if n := x * y; n%2 == 0 {
		f.Printf("n(%d) is even\n", n)
	} else {
		f.Printf("n(%d) is odd\n", n)
	}
	// f.Printf("n(%d) is undefined\n", n) // スコープ外なので n は未定義

}
