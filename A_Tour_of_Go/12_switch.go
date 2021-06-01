package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin" :
		fmt.Println("OS X.")
	case "linux" :
		fmt.Println("Linux.")
	default :
		fmt.Printf("%s", os)
	}

	/* 評価値の省略if~elseif~elseを避けるために使われたりする */
	switch t := time.Now(); {
	case t.Hour() < 12 :
		fmt.Println("おはよう")
	case t.Hour() <17 :
		fmt.Println("こんにちは")
	default :
		fmt.Println("こんばんは")
	}
}
/*
Go runs onLinux.
こんばんは
*/