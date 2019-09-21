package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// path:          httpやftpなどの論理パスを扱うパッケージ
	// path/filepath: OSファイルシステム上の物理パスを扱うパッケージ
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// httpリクエストは論理パスなのでpathパッケージを使う
		if ok, err := path.Match("/data/*.html", r.URL.Path); err != nil || !ok {
			http.NotFound(w, r)
			return
		}

		// 以降はパスを物理パスとして扱うのでpath/filepathパッケージを使う
		// filepath.Baseをpatn.Baseに置き換えて実行するとWindows環境では本来見えてはいけないソースコードファイルが露出してしまう
		// :8080/data/..\main.go
		name := filepath.Join(cwd, "data", filepath.Base(r.URL.Path))
		f, err := os.Open(name)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		defer f.Close()
		io.Copy(w, f)
	})

	http.ListenAndServe(":8080", nil)
}
