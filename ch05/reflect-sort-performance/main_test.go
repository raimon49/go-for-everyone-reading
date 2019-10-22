package main

import (
	"math/rand"
	"sort"
	"testing"
)

func randlist(n int) []int {
	// n個の配列を作成
	l := make([]int, n)
	for i := range l {
		l[i] = i
	}

	// 適当にシャッフル
	for i := range l {
		j := rand.Intn(i + 1)
		l[i], l[j] = l[j], l[i]
	}

	return l
}

func BenchmarkSortReflect(b *testing.B) {
	master := randlist(25)
	l := make([]int, 25)

	for i := 0; i < b.N; i++ {
		copy(l, master) // マスターから作業用バッファにデータをコピー

		// 自分で定義したSortwrapでリフレクションを利用したカスタムソート
		sort.Sort(NewSortwrap(l , func(i, j int) bool {
			return l[i] < l[j]
		}))
	}
}

func BenchmarkSortStandard(b *testing.B) {
	master := randlist(25)
	l := make([]int, 25)

	for i := 0; i < b.N; i++ {
		copy(l, master) // マスターから作業用バッファにデータをコピー

		// 組み込みのsort.Ints()でソート
		sort.Ints(l)
	}
}
