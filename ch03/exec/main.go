package main

import (
	"fmt"
	"os"
	"os/exec"
	"io"
	"log"
	"sync"
	"syscall"
)

func main() {
	tr(os.Stdin, os.Stdout, os.Stderr)

	// UNIX系環境ではsh -c経由で1つの引数としてコマンド + オプションを実行可能
	out, _ := exec.Command("sh", "-c", "ls -l").Output()
	fmt.Println(string(out))
}

func tr(src io.Reader, dst io.Writer, errDst io.Writer) error {
	// 実行するコマンドtr a-z A-Z
	cmd := exec.Command("tr", "a-z", "A-Z")
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	// コマンドの実行を開始
	var err error
	err = cmd.Start()
	if err != nil {
		return err
	}

	// 書き込みと読み込みを並行して行うための待ち受けグループ
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		// コマンドの標準入力にsrcからコピー
		_, err := io.Copy(stdin, src)
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.EPIPE {
			// ignore EPIPE
		} else if err != nil {
			log.Println("failed to write to STDIN", err)
		}

		stdin.Close()
		wg.Done()
	}()
	go func() {
		// コマンドの標準出力をdstにコピー
		io.Copy(dst, stdout)
		stdout.Close()
		wg.Done()
	}()
	go func() {
		// コマンドの標準エラー出力をerrDstにコピー
		io.Copy(errDst, stderr)
		stderr.Close()
		wg.Done()
	}()

	// 標準入出力のI/Oを行うgoroutineがすべて終わるまで待つ
	wg.Wait()

	// コマンドの終了を待つ
	return cmd.Wait()
}
