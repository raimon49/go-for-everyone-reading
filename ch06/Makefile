.DEFAULT_GOAL:= test

# Stringと名前のつくベンチマーク関数を全て実行
.PHONY: bench
bench:
	go test cat.go cat_test.go -bench String

# 実装コードとテストコードを一緒に引数に指定
.PHONY: test
test:
	go test -v calc.go calc_test.go
	go test -v hello.go hello_test.go
	go test -v unordered_test.go

.PHONY: reflect
reflect:
	go test -v reflect_test.go

.PHONY: race
race:
	cd data-race; make race

.PHONY: fixed-data-race
fixed-data-race:
	cd fixed-data-race; make race

.PHONY: setup
setup:
	go test -v setup_test.go

.PHONY: integration
integration:
	cd build-constrains; make

.PHONY: gists
gists:
	go test -v gists.go gists_test.go

.PHONY: gists_interface
gists_interface:
	go test -v gists_interface.go gists_interface_test.go

.PHONY: diff_gists
diff_gists:
	diff -u gists.go gists_interface.go           || exit 0
	diff -u gists_test.go gists_interface_test.go || exit 0
