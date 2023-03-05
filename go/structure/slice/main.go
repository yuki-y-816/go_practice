package slice

import (
	f "fmt"
)

type Point struct {
	X, Y int
}

type Points []*Point

func (p Points) toString() string {
	str := ""
	for _, p := range p {
		if str != "" {
			str += ","
		}

		if p == nil {
			str += "<nil>"
		} else {
			str += f.Sprintf("[%d, %d]", p.X, p.Y)
		}
	}

	return str
}

func Main() {
	f.Println("--- structure/slice ---")
	ps := make([]Point, 5)
	for i, v := range ps {
		f.Printf("ps[%d] : X=%d, Y=%d\n", i, v.X, v.Y)
	}

	ps2 := Points{}
	ps2 = append(ps2, &Point{X: 10, Y: 20})
	ps2 = append(ps2, nil)
	ps2 = append(ps2, &Point{X: 3, Y: 8})
	f.Printf("ps2.toString()=%s\n", ps2.toString())
}
