package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	tmp, _ := ioutil.TempFile(os.TempDir(), "tmp")
	defer tmp.Close()

	hash := sha256.New()

	// 両方に書き込むためのio.MultiWriter
	w := io.MultiWriter(tmp, hash)

	// io.Copyで標準入力からMultiWriterへコピー
	written, _ := io.Copy(w, os.Stdin)

	fmt.Printf("Wrote %d bytes to %s\nSHA256: %x\n",
		written,
		tmp.Name(),
		hash.Sum(nil),
	)
}
