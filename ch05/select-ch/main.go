package main

import (
	"os"
)

func main() {
	filename := "target.txt"
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	// ファイルから読み込んだバイト列を逐次chにする
	ch := make(chan []byte)
	go func() {
		defer close(ch)
		buf := make([]byte, 4096)
		for {
			n, err := f.Read(buf)
			if err != nil {
				return
			}
			ch <- buf[:n]
		}
	}()

	// 順次チャンネルからバイト列を読み込み、それを表示
	for in := range ch {
		os.Stdout.Write(in)
	}
}
