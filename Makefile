PHONY: run clean build

run:
	go run main.go

clean:
	rm -rf ./app

build: clean
	go build -o app && ./app