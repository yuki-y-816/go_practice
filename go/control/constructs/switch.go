package constructs

import f "fmt"

func numSwitch(i int) {
	switch i {
	case 1, 2:
		f.Println("1 or 2")
		fallthrough
	case 3, 4:
		f.Println("3 or 4")
	// case "5": // 型が違うからコンパイルエラー
	// f.Println("5")
	default:
		f.Println("unknown")
	}
}

func evenOrOdd(i int) {
	switch {
	case i%2 == 0:
		f.Println(i, "is even")
	case i%2 == 1:
		f.Println(i, "is odd")
	}
}

func chkType(x interface{}) {
	if x == nil {
		f.Println("this is nil")
	} else if i, isInt := x.(int); isInt {
		f.Println(i, "is int")
	} else if s, isString := x.(string); isString {
		f.Printf("%v is string\n", s)
	}
}

func chkType2(x interface{}) {
    switch x.(type) {
    case bool:
        f.Println("boolean")
    case int, uint:
        f.Println("integer or unsigned integer")
    case string:
        f.Println("string")
    default:
        f.Println("don't know")
    }
}

func chkType3(x interface{}) {
    switch v := x.(type) {
    case bool:
        f.Printf("%v is boolean\n", v)
    case int, uint:
        f.Printf("%v is integer or unsigned integer\n", v)
    case string:
        f.Printf("%v is string\n", v)
    default:
        f.Println("don't know")
    }
}

func SwitchFunc() {
	numSwitch(3)
	numSwitch(1)
	evenOrOdd(77)

	var x interface{} = 3.14
	//f.Println("x.(int) is", x.(int)) // interface{}によって隠蔽されてた元の型とは違う
	f.Println("x.(float64) is", x.(float64))
	i, isInt := x.(int)
	f.Println("x.(int) is", i, "isInt is", isInt)
	fl, isFloat64 := x.(float64)
	f.Println("x.(float64) is", fl, "isFloat64 is", isFloat64)

	chkType(nil)
	chkType(100)
	chkType("hello")

    chkType2(true)
    chkType2("hello!!")
    chkType2(x)

    chkType3(true)
    chkType3(7)
    chkType3("hello!!")
}
