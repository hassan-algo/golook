build:
	go build -o ./bin/execute ./main.go

run: build
	./bin/execute
