package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"os"
)

func main() {
	name := os.Args[1]
	s, _ := os.Stat(name)
	fmt.Printf(
		"%s: %s\n",
		name,
		humanize.Bytes(uint64(s.Size())),
	)
	fmt.Printf(
		"%s: %s\n",
		name,
		humanize.IBytes(uint64(s.Size())),
	)
	mb := "128MB"
	b, _ := humanize.ParseBytes(mb)
	fmt.Printf(
		"%s is %s bytes\n",
		mb,
		humanize.Comma(int64(b)),
	)
}
