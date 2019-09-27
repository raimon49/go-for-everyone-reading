package main

import (
	"reflect"
	"fmt"
)

type Point struct {
	X int
	Y int
}

// Package reflect: https://golang.org/pkg/reflect/
func main() {
	p := &Point{X: 10, Y: 5}

	// ValueOf()で型情報、格納されている値の取得
	rv := reflect.ValueOf(p)
	fmt.Printf("rv.Type = %v\n", rv.Type())
	fmt.Printf("rv.Kind = %v\n", rv.Kind())
	fmt.Printf("rv.Interface = %v\n", rv.Interface())

	// Point.Xの要素を取り出して100にセットし直す
	// 書籍ではこうなっているがpanicが発生する
	// xv := rv.Field(0)
	xv := reflect.Indirect(rv).Field(0)
	fmt.Printf("xv = %d\n", xv.Int())
	xv.SetInt(100)
	fmt.Printf("xv (after SetInt()) = %d\n", xv.Int())

	// ValueOf()で取り出したreflect.Valueは、それぞれの型に適合するメソッド以外を呼び出すとpanicを起こす
	rv1 := reflect.ValueOf(1)
	rv2 := reflect.ValueOf("Hello, World")
	rv3 := reflect.ValueOf([]byte{0xa,0xd})
	rv4 := reflect.ValueOf(make(chan string))
	rv1.Int()        // .String()だとpanic
	rv2.String()     // .Int()だとpanic
	rv3.Slice(0, 0)  // .MapKeys()だとpanic
	rv4.TrySend(rv2) // TrySend(rv1)だとpanic
}
