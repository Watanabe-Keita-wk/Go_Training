package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1 : 4]
	fmt.Println(s)
	/* [3 5 7] */

	fmt.Println(len(s))
	/* 3 */

	fmt.Println(cap(s))
	/* 5 */

	s[0] = 30
	fmt.Println(s, primes)
	/* [30 5 7] [2 30 5 7 11 13] */

	r := []bool{true, false, true}
	fmt.Println(r)
	/* [true false true] */

	fmt.Println(primes[:3])
	fmt.Println(primes[3:])
	/*
	[2 30 5]
	[7 11 13]
	*/

	a := make([]int, 3)
	b := make([]int, 0, 3)
	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
	/*
	3 3
	0 3
	*/

	/* スライスに追加 */
	a = append(a, 1)
	fmt.Println(a)
	/* [0 0 0 1] */
}