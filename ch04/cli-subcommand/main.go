package main


import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

// サブコマンドのインタフェース定義
type Command interface {
	// 簡単なコマンドの説明 ( todo -h )
	Synopsis() string

	// 詳細なヘルプメッセージ（ todo -h add )
	Help() string

	// 実際のコマンドの機能を記述
	// 引数（[]string）を受け取り終了ステータスを返す
	Run(args []string) int
}

// addコマンドの定義
type AddCommand struct {
}

func (c *AddCommand) Synopsis() string {
	return "Add todo task to list"
}

func (c *AddCommand) Help() string {
	return "Usage: todo add [option]"
}

// todo add xxx実行時の処理
func (c *AddCommand) Run(args []string) int {
	return 0
}

func main() {
	c := cli.NewCLI("todo", "0.1.0")

	// ユーザーの引数をCLIに登録
	c.Args = os.Args[1:]

	// サブコマンドを登録
	c.Commands = map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &AddCommand{}, nil
		},
	}

	// コマンドを実行
	exitCode, err := c.Run()
	if err != nil {
		fmt.Printf("Failed to execute: %s\n", err.Error())
	}

	// サブコマンドの終了ステータスに基づきコマンドを終了
	os.Exit(exitCode)
}
