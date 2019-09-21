package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// パス区切り文字列 "/" を直に書いて結合しない
	dir := filepath.Join(u.HomeDir, ".config", ".golangdir")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("created directory: %s", dir)
}
