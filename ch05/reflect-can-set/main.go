package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	X int
	Y int
}

func main() {
	p := Point{X: 10, Y: 5}
	rv := reflect.ValueOf(p)
	if !rv.Field(0).CanSet() {
		// 値渡しされているためrv.Field(0).SetInt()を呼ぶとpanicが起きるようGoは設計されている
		fmt.Println("Can't set value")
	}

	rvptr := reflect.ValueOf(&p)
	// 書籍のコードではpanicを起こすためreflect.Indirect()で括る
	// if f := reflect.rvptr.Field(0); f.CanSet() {
	if f := reflect.Indirect(rvptr).Field(0); f.CanSet() {
		// ポインタ渡しされ値のコピーでなく変数pそのものを参照しているためSetInt()が呼べる
		f.SetInt(0)
		fmt.Println("Done set value")
	}

}
