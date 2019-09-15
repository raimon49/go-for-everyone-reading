package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 取り扱うシグナルの決定
	trapSignals := []os.Signal {
		syscall.SIGHUP,  //  1: ハングアップ
		syscall.SIGINT,  //  2: インタラプト（割り込み）
		syscall.SIGTERM, // 15: ターミネート（killコマンドからの強制終了）
		syscall.SIGQUIT, //  3: 終了とコアダンプ
	}

	// シグナルを受信するチャンネルを定義
	sigCh := make(chan os.Signal, 1)
	// 受信するよう登録
	signal.Notify(sigCh, trapSignals...)
	// 別goroutineで処理を行う
	go doMain()
	// シグナルを受信するまでブロック
	sig := <-sigCh
	// 受信したシグナルが変数sigに渡されdoMainのループ処理が中断される
	fmt.Println("\nGot signal", sig)
}

func doMain() {
	// このdefer文は実行されない
	defer fmt.Println("done infintie loop")
	for {
		time.Sleep(1 * time.Second)
	}
}
