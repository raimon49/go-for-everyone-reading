.DEFAULT_GOAL:= build
.PHONY: build
build: clean setup
	go build .

.PHONY: setup
.ONESHELL:
setup:
	mkdir -p $(HOME)/.config/go-for-everyone-myapp
	cat << EOF > ~/.config/go-for-everyone-myapp/config.json
	{"Addr":"8080","Target":[{"Name":"foo","Threshold":3},{"Name":"bar","Threshold":4}]}
	EOF

.PHONY: clean
clean:
	rm -rf $(HOME)/.config/go-for-everyone-myapp
