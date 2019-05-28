package calc

import "testing"

// 加算のテストをする
// go test実行時にテストの対象となるよう名前をTextXXXXにして、
// 引数に*testing.Tを渡す
func TestSum(t *testing.T) {
	if sum(1, 2) != 3 {
		// t.Fatalでテストが失敗したことを返す
		t.Fatal("sum(1, 2) should be 3, but doesn't match")
	}
}
