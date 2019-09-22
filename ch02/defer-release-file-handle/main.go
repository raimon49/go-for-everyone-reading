package main

import (
	"os"
)

func doSomething() error {
	err := os.Mkdir("newdir", 0755)
	if err != nil {
		return err
	}

	// 2) 次にディレクトリが削除される
	defer os.RemoveAll("newdir")

	f, err := os.Create("newdir/newfile")
	if err != nil {
		return err
	}

	// deferは呼びされたスコープの逆順で実行される
	// 後処理が必要なオブジェクトを作った際にはエラーチェックと直後にdeferで閉じる慣習を持つとよい
	// 1) 最初にファイルハンドルが閉じられる
	defer f.Close()

	return nil
}

func main() {
	doSomething()
}
