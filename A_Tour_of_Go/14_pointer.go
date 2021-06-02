package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(p, *p)
	n := p
	fmt.Println(*n)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
/*
0xc000100010 42
42
73
*/
