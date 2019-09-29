package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Point struct {
	X int `mytag:"x,omitempty"`
	Y int `mytag:"y,omitempty"`
}

func Marshal(v interface{}) ([]byte, error) {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Map:
		// 引数vの型がmapのとき
		fmt.Println("type of arguments 'v' is 'map'")
	case reflect.Struct:
		// 引数vの型がstructのとき
		fmt.Println("type of arguments 'v' is 'struct'")
	default:
		// map/struct以外ではエラー
		return nil, errors.New("Marshal: unsupported type(" + rv.Type().String() + ")")
	}

	return nil, nil
}

func main() {
	Marshal(Point{X: 10, Y:0})
}
