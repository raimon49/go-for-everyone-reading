package main


import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)


func TestRun(t *testing.T) {
	// テスト中のRun関数の出力先はメモリバッファにする
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	// os.Argsをテストしたいコマンド実行の文字列から作成
	args := strings.Split("gobook -version", " ")

	cli.Run(args)

	expected := fmt.Sprintf("gobook version v0.1.0")
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("expected %q to %q", outStream.String(), expected)
	}
}
