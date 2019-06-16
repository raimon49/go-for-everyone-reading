package deepequal

import (
	"reflect"
	"testing"
)

type T struct {
	x int
	ss []string
	m map[string]int
}

func TestStruct(t *testing.T) {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	t1 := T{
		x: 1,
		ss: []string{"a", "b"},
		m: m1,
	}
	t2 := T{
		x: 1,
		ss: []string{"a", "b"},
		m: m1,
	}

	// reflect.DeepEqual()でstructの全てのフィールドを比較する
	if !reflect.DeepEqual(t1, t2) {
		// Goでは欲しかった期待値についての分かり易いエラーメッセージを自分で書く
		t.Errorf("want %#v got %#v", t1, t2)
	}
}
