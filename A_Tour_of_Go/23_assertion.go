package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	/* 型アサーションとswitchの組み合わせ。直感で理解できなかったから調べた。
	　　typeがswitchの条件になる。型switchというらしい
	*/
	switch value := i.(type) {
	case string:
		fmt.Printf("\"%s\"は文字列型です", value)
	case int:
		fmt.Printf("%d は整数型です", value)
	case float64:
		fmt.Printf("%f は浮動小数点型です", value)
	default:
		fmt.Println(value, "は知らない型です")
	}
}
/*
hello
hello true
0 false
"hello"は文字列型です
*/