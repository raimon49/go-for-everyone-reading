package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
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

		// keyはreflect.Value型のキーを格納しており、これを使って変数rvからキーに対応する型を取り出せる
		for _, key := range rv.MapKeys() {
			mv := rv.MapIndex(key)
			fmt.Println("mapValue is " + mv.Type().String())
		}
	case reflect.Struct:
		// 引数vの型がstructのとき
		fmt.Println("type of arguments 'v' is 'struct'")

		rt := rv.Type()
		for i := 0; i < rt.NumField(); i++ {
			// reflect.StructField型
			ftv := rt.Field(i)
			// reflect.Value型
			fv := rv.Field(i)
			fmt.Println("Field " + ftv.Name + " is " + fv.Type().String())

			if ftv.PkgPath != "" {
				continue
			}

			tag := ftv.Tag.Get("mytag")
			parts := strings.Split(tag, ",")
			shortname := parts[0]

			fmt.Println("Field " + ftv.Name + "'s shortname is " + shortname + " defined in the 'struct tag'")
		}
	default:
		// map/struct以外ではエラー
		return nil, errors.New("Marshal: unsupported type(" + rv.Type().String() + ")")
	}

	return nil, nil
}

func main() {
	Marshal(Point{X: 10, Y:0})
}
