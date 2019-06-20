package main

import (
	"fmt"
	"sync"
)

func main() {
	// sync.Mutexやsync.RWMutexを使ってスレッドセーフな実装にする
	var mux sync.Mutex
	c := make(chan bool)
	m:= make(map[string]string)
	go func() {
		mux.Lock()
		m["1"] = "a"
		mux.Unlock()
		c <- true
	}()

	mux.Lock()
	m["2"] = "b"
	mux.Unlock()

	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
