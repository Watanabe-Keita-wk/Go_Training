package main

import "fmt"

func main() {
	janken := []string{"グー", "チョキ", "パー"}

	for me := 0; me < len(janken); me++ {
		fmt.Printf("\n私が%sのとき：\n", janken[me])
		for you := 0; you < len(janken); you++ {
			fmt.Printf("あなたが%sなら", janken[you])
			//↓この考え{(3 + me - you) % 3}思いついたのすごいな、、、
			switch (3 + me - you) % 3 {
				case 0:
					fmt.Printf("あいこ\n")
				case 1:
					fmt.Printf("勝ち\n")
				case 2:
					fmt.Printf("負け\n")
			}
		}

	}
}
/*

私がグーのとき：
あなたがグーならあいこ
あなたがチョキなら負け
あなたがパーなら勝ち

私がチョキのとき：
あなたがグーなら勝ち
あなたがチョキならあいこ
あなたがパーなら負け

私がパーのとき：
あなたがグーなら負け
あなたがチョキなら勝ち
あなたがパーならあいこ
*/