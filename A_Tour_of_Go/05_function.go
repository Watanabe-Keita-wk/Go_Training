package main

import (
	"fmt"
)

/* 返り値の型指定 */
func add(x int, y int) int {
	return x + y
}

/* 引数の型宣言省略 */
func swap(x , y string) (string, string) {
	return y, x
}

/* 名前付き返り値を返す */
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

func main() {
	var sum int = add(24, 48)
	fmt.Println(sum)
	fmt.Println(swap("World!", "Hello"))
	fmt.Println(split(sum))
	/*
	72
	Hello World!
	32 40
	*/
}