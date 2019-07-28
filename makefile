phony: run build

build:
	go build -o ./bin/exchange service/exchange/main.go

run: build
	./bin/exchange