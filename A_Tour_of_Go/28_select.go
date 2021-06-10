package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Tick(n): n秒ごとの実行プロセスを作成
	tick := time.Tick(100 * time.Millisecond) // 0.1秒ごとの処理

	// time.After(n): n秒後の実行プロセスを作成
	boom := time.After(500 * time.Millisecond) // 0.5秒後の処理

	// 無限ループ
	//selectの仕組みはわかったけど、実行ぷ実行プロセスの作成がよくわからん。、、
	for {
		select {
		case <-tick:
			// 0.1秒ごとに "tick." を出力
			fmt.Println("tick.")
		case <-boom:
			// 0.5秒後に "BOOM!" を出力してループ終了
			fmt.Println("BOOM!")
			return
		default:
			// 上記のいずれのチャネルも準備されていない場合は "    ." を出力して 0.05秒待機
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
/*
   	.
    .
tick.
    .
    .
tick.
    .
    .
tick.
    .
    .
tick.
    .
    .
tick.
BOOM!
*/