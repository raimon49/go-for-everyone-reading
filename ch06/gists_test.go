package gists

import (
	"testing"
	"io"
	"strings"
)

func TestGetGists(t *testing.T) {
	// 実装側のdoGistsRequestを上書き
	// このdoGistsRequestはTestGetGistsのスコープの中でのみ有効
	doGistsRequest = func(user string) (io.Reader, error) {
		return strings.NewReader (`
[
	{"html_url": "https://gist.github.com/example1"},
	{"html_url": "https://gist.github.com/example2"}
]`), nil
	}

	urls, err := ListGists("test")
	if err != nil {
		t.Fatalf("list gists caused error: %s", err)
	}

	if expected := 2; len(urls) != expected {
		t.Fatalf("want %d, got %d", expected, len(urls))
	}
}
