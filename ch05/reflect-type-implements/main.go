package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func ReadData(rdr io.Reader) {
}

func main() {
	v := os.Stdin
	rv := reflect.ValueOf(v)

	var rdr io.Reader
	// reflect.TypeOf(rdr)だとnilなTypeが返り、Implements()はnilを受け付けずpanicが起きるため*io.Reader→io.Readerで取得
	ptrT := reflect.TypeOf(&rdr)
	T := ptrT.Elem()

	if rv.Type().Implements(T) {
		fmt.Println("os.Stdin is implemented io.Reader.")
	}

	// ポインタ経由で型Tを取得するワンライナー
	T = reflect.TypeOf((*io.Reader)(nil)).Elem()
	if rv.Type().Implements(T) {
		fmt.Println("os.Stdin is implemented io.Reader. (via oneliner)")
	}
}
