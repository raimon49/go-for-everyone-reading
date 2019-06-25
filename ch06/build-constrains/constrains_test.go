// +build integration

package foo

import (
	"fmt"
	"testing"
)

func TestSomothing(t *testing.T) {
	fmt.Println("Run only when '-tags=integration is specified'")
}
