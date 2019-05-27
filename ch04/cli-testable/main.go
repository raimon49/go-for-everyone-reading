package main


import (
	"fmt"
	"io"
	"os"
)

// 標準出力と標準エラーのストリームをフィールドとして持つCLIの定義
type CLI struct {
	outStream, errStream io.Writer
}


func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}

// 引数処理を含めた具体的な処理
func (c *CLI) Run(args []string) int {
	var version = "v0.1.0"
	fmt.Fprintf(c.errStream, "gobook version %s \n", version)

	return 1
}
