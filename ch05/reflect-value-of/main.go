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
	// それぞれの型と違うメソッドを呼ぶとpanicが起きるため、検査してから取り出す
	rv1 := reflect.ValueOf(1)
	if rv1.Kind() == reflect.Int {
		fmt.Printf("rv1 type is Int: %d\n", rv1.Int())
	}

	rv2 := reflect.ValueOf("Hello, World")
	if rv2.Kind() == reflect.String {
		fmt.Printf("rv2 type is String: %s\n", rv2.String())
	}

	rv3 := reflect.ValueOf([]byte{0xa,0xd})
	if rv3.Kind() == reflect.Slice {
		fmt.Println("rv3 type is Slice")
		rv3.Slice(0, 0)
	}

	rv4 := reflect.ValueOf(make(chan string))
	if rv4.Kind() == reflect.Chan {
		fmt.Println("rv3 type is Chan")
		rv4.TrySend(rv2) // TrySend(rv1)だとpanic
	}
}
