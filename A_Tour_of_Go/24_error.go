package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "it didn't work"}
}

func main() {
	if err := run(); err != nil {
		// エラーが起きた場合、そのエラー値を出力
		fmt.Println(err)
	}
}
/*
at 2009-11-10 23:00:00 +0000 UTC m=+0.000000001, it didn't work
*/