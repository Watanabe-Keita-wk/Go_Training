package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	v.X = 4
	fmt.Println(v)

	p := &v

	(*p).X = 1e3
	p.Y = 1e2
	fmt.Println(v)

	var (
		v1 = Vertex{X: 1}
		v2 = Vertex{}
		pv = &Vertex{}
	)
	fmt.Println(v1, v2, pv)
}
/*
{1 2}
{4 2}
{1000 100}
{1 0} {0 0} &{0 0}
*/