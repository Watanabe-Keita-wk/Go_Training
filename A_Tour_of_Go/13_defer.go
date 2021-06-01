package main

import "fmt"

func hello() {
	/* func hello()が終わった後に実行される */
	defer fmt.Println("world")
	fmt.Print("Hello,")
}

/* 複数の処理が実行される場合、後入れ先だしで実行される。*/
func count() {
	for i := 0; i < 11; i++ {
		if i == 0 {
		defer fmt.Println(i)
		} else {
			defer fmt.Print(i, ",")
		}
	}
	fmt.Println("counting")
}

func main() {
	hello()
	count()
}
/*
Hello,world
counting
10,9,8,7,6,5,4,3,2,1,0
*/