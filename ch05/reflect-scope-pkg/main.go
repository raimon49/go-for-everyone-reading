package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	p    int // 小文字で始まる要素はパッケージ外からreflectでアクセスできない
}

func main() {
	p := Person{}
	rt := reflect.TypeOf(p)
	rv := reflect.ValueOf(p)

	for i := 0; i < rv.NumField(); i++ {
		ft := rt.Field(i)
		fv := rv.Field(i)

		if ft.PkgPath == "" {
			fmt.Printf("ft(i) -> %#v\n", ft)
			fmt.Printf("fv(i) -> %#v\n", fv.Interface())
		} else {
			fmt.Printf("fv(i) PkgPath: %s", ft.PkgPath)
		}
	}
}
