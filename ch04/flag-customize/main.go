package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	ExitCodeOK        int = iota
	ExitCodeError
	ExitCodeFileError
)

func main() {
	flags := flag.NewFlagSet("example", flag.ContinueOnError)
	// メッセージの出力先を標準エラーから標準出力に変更
	flags.SetOutput(os.Stdout)
	// ヘルプメッセージの変更
	flags.Usage = func() {
		fmt.Println("Usage: This is customized usage message")
	}
	if err := flags.Parse(os.Args[1:]); err != nil {
		// flag.ExitOnErrorでなくflag.ContinueOnErrorで独自のエラー処理
		fmt.Printf("Handle parse errors")
		os.Exit(ExitCodeError)
	}

	os.Exit(ExitCodeOK)
}
