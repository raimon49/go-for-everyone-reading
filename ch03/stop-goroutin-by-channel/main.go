package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	queue := make(chan string)
	for i := 0; i < 2; i++ {
		// 2つのワーカーを生成
		wg.Add(1)
		go fetchURL(queue)
	}

	queue <- "http://www.example.com"
	queue <- "http://www.example.net"
	queue <- "http://www.example.net/foo"
	queue <- "http://www.example.net/bar"

	close(queue) // goroutineに終了を伝える
	wg.Wait()    // 全てのgoroutineが終了するまで待つ
}

func fetchURL(queue chan string) {
	for {
		url, more := <- queue // closeされるとmoreはfalseになる
		if more {
			// URL取得処理（省略）
			fmt.Println("fetching", url)
		} else {
			fmt.Println("worker exit")
			wg.Done()
			return
		}
	}
}

/*
 * チャネルから読みだした2つ目の値がfalseか調べる代わりにrangeでも同じことが可能
 * チャネルがcloseされるとrange queueのforループを抜けるため、goroutineを終了できる
func fetchURL(queue chan string, done chan bool) {
	for url := range queue {
		// URL取得処理（省略）
		fmt.Println("fetching", url)
	}

	fmt.Println("worker exit")
	done <- true
}
*/
