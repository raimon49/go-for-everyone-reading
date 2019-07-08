package gists

import (
	"bytes"
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

type Gist struct {
	Rawurl string `"json:html_url"`
}

// doGistsRequestはグローバルスコープに定義された関数オブジェクト（テストコード側から書き換え可能）
var doGistsRequest = func(user string) (io.Reader, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/gists", user))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return nil, err
	}

	return &buf, nil
}

func ListGists(user string) ([]string, error) {
	r, err := doGistsRequest(user)
	if err != nil {
		return nil, err
	}

	var gists []Gist
	if err := json.NewDecoder(r).Decode(&gists); err != nil {
		return nil, err
	}

	urls := make([]string, 0, len(gists))
	for _, u := range gists {
		urls = append(urls, u.Rawurl)
	}

	return urls, nil
}
