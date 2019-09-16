package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type MySignal struct {
	message string
}

func (s MySignal) String() string {
	return s.message
}

func (s MySignal) Signal() {}

func main() {
	log.Println("[info] Start")
	trapSignals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	}

	// シグナルを受信するチャンネルを定義
	sigCh := make(chan os.Signal, 1)

	// 受信するよう登録
	signal.Notify(sigCh, trapSignals...)

	// 10秒後にsigChにMySignalの値を送信
	time.AfterFunc(10*time.Second, func() {
		sigCh <- MySignal{"timed out"}
	})

	// シグナルを受信するまで待ち受ける
	sig := <-sigCh

	switch s := sig.(type) { // 型アサーションで受信したシグナルを判別
	case syscall.Signal:
		// osからのシグナル受信の場合
		log.Printf("\n[info] Got signal: %s(%d)", s, s)
	case MySignal:
		// アプリケーション独自定義したシグナル受信の場合
		log.Printf("\n[info] %s", s) // .String() が評価される
	}
}
