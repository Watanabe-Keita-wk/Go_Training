package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64}

func main() {
	//foreach的な？
	for i, v := range pow {
		fmt.Printf("2^%d = %d ", i, v)
	}

	for _, v := range pow {
		fmt.Println(v)
	}
}
/*
2^0 = 1 2^1 = 2 2^2 = 4 2^3 = 8 2^4 = 16 2^5 = 32 2^6 = 64 1
2
4
8
16
32
64
*/