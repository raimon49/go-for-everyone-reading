package main

import (
	"reflect"
	"time"
	"fmt"
)

var timeT = reflect.TypeOf(time.Time{})

func MakeTime() *time.Time {
	rv := reflect.New(timeT)
	return rv.Interface().(*time.Time)
}

func main() {
	// t := time.Time{}と同値
	t := MakeTime()
	year, month, day := t.Date()
	fmt.Printf("year = %v\n", year)
	fmt.Printf("month = %v\n", month)
	fmt.Printf("day = %v\n", day)

	// structの動的な定義と作成
	rt := reflect.StructOf([]reflect.StructField{
		{Name: "Number", Type: reflect.TypeOf(0)}, // 整数型フィールド
		{Name: "Str", Type: reflect.TypeOf("")},   // 文字列型フィールド
	})

	rv := reflect.New(rt)
	fmt.Printf("dynamic created struct: %#v", rv)
}
