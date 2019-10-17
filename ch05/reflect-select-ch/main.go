package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

func makeChannelsForFiles(files []string) ([]reflect.Value, error) {
	cs := make([]reflect.Value, len(files))

	for i, fn := range files {
		// データを流すためのチャンネルを作成
		ch := make(chan []byte)

		// ファイルをオープン
		f, err := os.Open(fn)
		if err != nil {
			return nil, err
		}

		go readFromFile(ch, f)

		cs[i] = reflect.ValueOf(ch)
	}

	return cs, nil
}

func readFromFile(ch chan []byte, f *os.File) {
	defer close(ch) // 全て終わったら引数のチャンネルを閉じる
	defer f.Close() // 全て終わったら引数のファイルを閉じる

	buf := make([]byte, 4096)
	for {
		// 読み込めるデータがあればそれをチャンネルに渡す
		// ここではエラーがあっても敢えて処理を続ける（io.EOFを受け取ったらtailできなくなってしまうため）
		if n, err := f.Read(buf); err == nil {
			ch <- buf[:n]
		}
	}
}

func makeSelectCases(cs ...reflect.Value) ([]reflect.SelectCase, error) {
	// 与えられた分のchanの数だけSelectCaseを作成
	cases := make([]reflect.SelectCase, len(cs))

	// TODO

	return cases, nil
}

func doSelect(cases []reflect.SelectCase) {
	// TODO
}

func main() {
	if err := _main(); err != nil {
		fmt.Println("%s\n", err)
		os.Exit(1)
	}
}

func _main() error {
	if len(os.Args) < 2 {
		return errors.New("prog [file1 file2 ...]")
	}

	// シグナル処理のおまじない
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// os.Argsの最初の引数はこのコマンド名で、2番目以降が対象ファイル名
	channels, err := makeChannelsForFiles(os.Args[1:])
	if err != nil {
		return nil
	}

	// 引数のファイル分だけループして作成されたチャンネルで動的にselect caseを作成
	cases, err := makeSelectCases(channels...)
	if err != nil {
		return err
	}

	go doSelect(cases)

	return nil
}
