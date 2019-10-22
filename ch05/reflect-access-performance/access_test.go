package main

import (
	"reflect"
	"testing"
)

type StructAccess struct {
	Int int
}

func BenchmarkDetectTypeReflect(b *testing.B) {
	var s interface{} = StructAccess{Int: 100}
	for i := 0; i < b.N; i++ {
		rv := reflect.ValueOf(s)
		if rv.Type().Name() == "StructAccess" {
			_ = s.(StructAccess).Int // Intフィールドへのリフレクション経由でのアクセス
		}
	}
}

func BenchmarkDetectNone(b *testing.B) {
	s := StructAccess{Int: 100}
	for i := 0; i < b.N; i++ {
		_ = s.Int; // 型が分かっている前提でIntフィールドにそのままアクセス
	}
}

func BenchmarkDetectTypeAssert(b *testing.B) {
	var s interface{} = StructAccess{Int: 100}
	for i := 0; i < b.N; i++ {
		if sa, ok := s.(StructAccess); ok {
			_ = sa.Int // リフレクションは扱わず型アサーションだけでアクセス
		}
	}
}
