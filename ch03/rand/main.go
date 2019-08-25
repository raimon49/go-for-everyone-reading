package main

import (
	"encoding/binary"
	crand "crypto/rand"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var s int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &s); err != nil {
		// crypto/randからReadできなかった時の代替として現在時刻（ナノ秒単位）を利用
		s = time.Now().UnixNano()
	}

	rand.Seed(s)
	// 0 <= n < 100となるintの乱数を生成
	n := rand.Intn(100)
	fmt.Println(n)
}
