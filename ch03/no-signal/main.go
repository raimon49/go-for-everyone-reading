package main

import (
	"fmt"
	"time"
)

func main() {
	// forループ実行中にCtrl + CでSIGINTを受信したら、deferは実行されず強制終了される（Goのデフォルト挙動）
	defer fmt.Println("done")

	for {
		// 単に無限ループする処理
		time.Sleep(1 * time.Second)
	}
}
