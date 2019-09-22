package main

import (
	"fmt"
	"encoding/json"
	"runtime"
	"os"
	"path/filepath"
)

type target struct {
	Name      string `json:'name'`
	Threshold int    `json: 'threshold'`
}

type config struct {
	Addr   string   `json:'addr'`
	Target []target `json:'target'`
}

func loadConfig() (*config, error) {
	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		// WindowsでHOME環境変数が設定されていない場合
		home = os.Getenv("APPDATA")
	}

	fname := filepath.Join(home, ".config", "go-for-everyone-myapp", "config.json")
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	var cfg config
	err = json.NewDecoder(f).Decode(&cfg)

	return &cfg, err
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	b, _ := json.MarshalIndent(&cfg, "", "  ")
	fmt.Println(string(b))
}
