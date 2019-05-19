package main

import (
	"flag"
	"fmt"
	"strings"
)

type strSliceValue []string

// Valueインタフェースの実装（error)
func (v *strSliceValue) Set(s string) error {
	strs := strings.Split(s, ",")
	*v = append(*v, strs...)

	return nil
}

// Valueインタフェースの実装（string)
func (v *strSliceValue) String() string {
	return strings.Join(([]string)(*v), ",")
}

func main() {
	// 独自のフラグを利用
	var spacies []string
	flag.Var((*strSliceValue)(&spacies), "spacies", "")

	flag.Parse()

	if len(spacies) > 0 {
		fmt.Println("spacies:")
		for _, str := range spacies {
			fmt.Println(str)
		}
	}
}
