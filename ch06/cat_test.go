package cat

import (
	"fmt"
	"testing"
)

// ベンチマーク用のトークンを作成
// 引数で指定された長さの文字列のスライスを作って返す
func seed(n int) []string {
	s := make([]string, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, "a")
	}

	return s
}

// ベンチマーク用のヘルパー
// テストしたい文字列と、文字列結合を行う関数手続きを渡し、ベンチマークを実行
func bench(b *testing.B, n int, f func(...string) string) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f(seed(n)...)
	}
}

func BenchmarkStringCat3(b *testing.B) { bench(b, 3, cat) }
func BenchmarkStringBuf3(b *testing.B) { bench(b, 3, buf) }

func BenchmarkStringCat100(b *testing.B) { bench(b, 100, cat) }
func BenchmarkStringBuf100(b *testing.B) { bench(b, 100, buf) }

func BenchmarkStringCat10000(b *testing.B) { bench(b, 10000, cat) }
func BenchmarkStringBuf10000(b *testing.B) { bench(b, 10000, buf) }

// サブベンチマークで1つのベンチマーク手続きの中で全てを実行
func BenchmarkStringConcatenate(b *testing.B) {
	benchCases := []struct {
		name string
		n    int
		f    func(...string) string
	}{
		{"Cat", 3, cat},
		{"Buf", 3, buf},
		{"Cat", 100, cat},
		{"Buf", 100, buf},
		{"Cat", 10000, cat},
		{"Buf", 10000, cat},
	}

	for _, c := range benchCases {
		b.Run(fmt.Sprintf("%s%d", c.name, c.n), func(b *testing.B) {
			bench(b, c.n, c.f)
		})
	}
}
