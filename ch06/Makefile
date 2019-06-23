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
