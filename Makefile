all:
	go run main.go

linux:
	GOOS=linux go build -o bin/main main.go

windows:
	GOOS=windows go build -o bin/main.exe main.go
