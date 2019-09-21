package main

import (
	"fmt"
	"sync"
	"context"
)

var wg sync.WaitGroup

func main() {
	// キャンセルするためのContextを生成
	ctx, cancel := context.WithCancel(context.Background())
	/* context.WithCancel()の変わりにcontext.withTimeout()を使うと
	 * 指定時間が経過後にctx.Done()から読み出し可能になるため、タイムアウトとキャンセルを統一的に扱える
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	 */
	queue := make(chan string)
	for i := 0; i < 2; i++ {
		// 2つのワーカーを生成
		wg.Add(1)
		go fetchURL(ctx, queue)
	}

	queue <- "http://www.example.com"
	queue <- "http://www.example.net"
	queue <- "http://www.example.net/foo"
	queue <- "http://www.example.net/bar"

	cancel()  // ctxを全て終了する
	wg.Wait() // 全てのgoroutineが終了するまで待つ
}

func fetchURL(ctx context.Context, queue chan string) {
	for {
		select {
		case <- ctx.Done():
			fmt.Println("worker exit")
			wg.Done()
			return
		case url := <-queue:
			// URL取得処理（省略）
			fmt.Println("fetching", url)
		}
	}
}
