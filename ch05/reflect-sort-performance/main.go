package main

import (
	"reflect"
)

type Sortwrap struct {
	value    reflect.Value       // 任意の配列が入ったreflect.Value
	lessfunc func(int, int) bool // 上記value内の値を比較するための関数
}

// コンストラクタ。配列と配列要素を比較する関数を受け取る
func NewSortwrap(v interface{}, lessfunc func(int, int) bool) *Sortwrap {
	return &Sortwrap{
		value:    reflect.ValueOf(v),
		lessfunc: lessfunc,
	}
}

// sort.Interfaceで定義されたLen(), Less(), Swap()を用意

// sort.InterfaceのLen()を満たすメソッド
func (s *Sortwrap) Len() int {
	return s.value.Len()
}

// sort.Interface()のLess()を満たすメソッド
func (s *Sortwrap) Less(i, j int) bool {
	// 型Sortwrapがコンストラクタで受け取った関数に処理を委譲
	return s.lessfunc(i, j)
}

// sort.InterfaceのSwap()を満たすメソッド
func (s *Sortwrap) Swap(i, j int) {
	value := s.value
	v1 := value.Index(i).Interface()
	v2 := value.Index(j).Interface()

	// reflectパッケージを経由して配列内の要素を入れ替え
	value.Index(j).Set(reflect.ValueOf(v1))
	value.Index(i).Set(reflect.ValueOf(v2))
}
