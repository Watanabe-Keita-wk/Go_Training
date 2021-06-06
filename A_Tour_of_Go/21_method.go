package main

import (
	"fmt"
	"math"
)

/*
Goにはクラスの概念がない。その代わりがstructっぽい。
そして、クラスにメソッドがあるようにstruct構造体にメソッドを定義できる。
*/
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

type Float float64

func (f Float) Abs() Float {
	if f < 0 {
		return -f
	}
	return f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	v.Scale(2)
	fmt.Println(v)

	f := Float(-0.1)
	fmt.Println(f.Abs())
}
/*
5
{3 4}
0.1
*/