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
			fmt.Printf("fv(i) PkgPath: %s\n", ft.PkgPath)
		}
	}

	val := reflect.TypeOf(Person{})
	fmt.Printf("Method Name = %s\n", val.Method(0).Name) // Foo

	val = reflect.TypeOf(&Person{})
	fmt.Printf("Method Name = %s\n", val.Method(1).Name) // PtrFoo

}

// レシーバがポインタ
func (p *Person) PtrFoo() string {
	return "ptr"
}

// レシーバが実体
func (p Person) Foo() string {
	return "val"
}
