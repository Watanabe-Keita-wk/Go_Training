package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/* printfで特定の型を文章の中に入れる。*/
func main() {
	fmt.Printf("Now you have %g probrems.\n", math.Sqrt(7))
	
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Your lucky number is %d.\n", rand.Intn(10))
	/*
	Now you have 2.6457513110645907 probrems.
	Your lucky number is [0〜9の乱数].
	*/
}