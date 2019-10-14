package main

import (
	"errors"
	"fmt"
	"os"
	// "os/signal"
	// "syscall"
)

func main() {
	if err := _main(); err != nil {
		fmt.Println("%s\n", err)
		os.Exit(1)
	}
}

func _main() error {
	if len(os.Args) < 2 {
		return errors.New("prog [file1 file2 ...]")
	}

	return nil
}
