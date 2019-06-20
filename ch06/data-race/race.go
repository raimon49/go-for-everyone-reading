package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	// 排他制御なしに並列にmapへアクセス
	m:= make(map[string]string)
	go func() {
		m["1"] = "a" // 競合する1つめのアクセス
		c <- true
	}()
	m["2"] = "b" // 競合する2つめのmain goroutineからのアクセス

	// cはunbuffered channelなのでchannelへのsendが完了されるまで待つ
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
