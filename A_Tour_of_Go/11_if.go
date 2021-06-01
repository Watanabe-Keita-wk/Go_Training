package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim1, lim2 float64) float64 {
	if v := math.Pow(x, n); v < lim1 {
		return v
	} else if v >= lim1 && v < lim2 {
		return lim1
	} else {
		return lim2
	}
}

func main() {
	fmt.Println(
		pow(3, 2, 18, 40),
		pow(3, 3, 18, 40),
		pow(3, 4, 18, 40),
	)
}
/* 9 18 40 */