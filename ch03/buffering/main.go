package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
)

func main() {
	var output io.Writer
	if isatty.IsTerminal(os.Stdout.Fd()) {
		// 標準出力が端末なら出力先はos.Stdoutそのもの
		output = os.Stdout
	} else {
		// 標準出力が端末でなければbufio.Writerでラップ
		output = bufio.NewWriter(os.Stdout)
	}

	for i := 0; i < 100; i++ {
		fmt.Fprintln(output, strings.Repeat("x", 100))
	}

	if _o, ok := output.(*bufio.Writer); ok {
		// bufio.Writerでは最後にFlush()が必要
		_o.Flush()
	}
}
