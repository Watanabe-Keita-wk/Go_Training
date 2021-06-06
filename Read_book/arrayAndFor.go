package main

import "fmt"

func main() {
	//GoではシングルクォーテーションはNG?
	items := []string{"鉛筆", "消しゴム", "ノート", "ペン", "付箋"}
	prices := []int{100, 50, 300, 150, 200}
	var sum int = 0

	for i := 0; i < len(items); i++ {
		sum += prices[i]
		fmt.Printf("%s：%d円 小計：%d円\n", items[i], prices[i],sum)
	}
}
/*
鉛筆：100円 小計：100円
消しゴム：50円 小計：150円
ノート：300円 小計：450円
ペン：150円 小計：600円
付箋：200円 小計：800円
*/