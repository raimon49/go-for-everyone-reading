.PHONY: test

.DEFAULT_GOAL:= test

bench:
	go test cat.go cat_test.go -bench String

test:
	go test -v calc.go calc_test.go
	go test -v hello.go hello_test.go
	go test -v unordered_test.go
