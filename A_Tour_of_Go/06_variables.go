package main
import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)

	var x, y int = 1, 2
	/* 型が明らかな場合は型宣言省略可 */
	var z = "ゼット"
	fmt.Println(x, y, z)
	/*
	0 false false false
	1 2 ゼット
	*/
}