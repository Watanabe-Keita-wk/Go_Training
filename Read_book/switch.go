package main

import "fmt"

func main() {
	age := []int{18, 22, 56, 39, 42, 62,  27}
	gen := [3]int{}

	/*
	switchの条件値にbool型(true)を置くことで、true or falseで条件分岐
	この形はよく使われるため、dummyは省略でswitch {}でできる。
	*/
	dummy := true
	for i := 0; i < len(age); i++ {
		switch dummy {
		case age[i] > 49:
			gen[2]++
		case age[i] > 29:
			gen[1]++
		default:
			gen[0]++
		}
	}

	fmt.Printf("%d名のうち、：\n", len(age))
	fmt.Printf("50代以上は%d名\n", gen[2])
	fmt.Printf("30~40代は%d名\n", gen[1])
	fmt.Printf("20代以下は%d名\n", gen[0])
}
/*
7名のうち、：
50代以上は2名
30~40代は2名
20代以下は3名
*/