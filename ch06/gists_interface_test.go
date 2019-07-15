package gists_interface

import (
	"testing"
	"io"
	"strings"
)

// Doerのインタフェースを満たすダミーのstruct
type dummyDoer struct {}

// doGistsRequestのダミー実装。HTTP Requestを送らないでダミーのデータを返す。
func (d *dummyDoer) doGistsRequest(user string) (io.Reader, error) {
	return strings.NewReader (`
[
	{"html_url": "https://gist.github.com/example1"},
	{"html_url": "https://gist.github.com/example2"}
]`), nil
}

func TestGetGists2(t *testing.T) {
	// dummyDoerはDoerの実装なので、Clientに渡すことができる。
	c := &Client{&dummyDoer{}}
	urls, err := c.ListGists("test")
	if err != nil {
		t.Fatalf("list gists caused error: %s", err)
	}

	if expected := 2; len(urls) != expected {
		t.Fatalf("want %d, got %d", expected, len(urls))
	}
}
