package setup

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// テスト用の環境セットアップなどの初期化処理
	setup()
	// テストコード本体の実行
	exitCode := m.Run()
	// テスト環境の片付け
	shutdown()
	// テストの終了
	os.Exit(exitCode)
}

func setup() {
	fmt.Println("Call setup")
}

func TestRun(t *testing.T) {
	fmt.Println("Run test code")
}

func shutdown() {
	fmt.Println("Call shutdown")
}
