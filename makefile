build:
	go build -o server ./cmd/api/main.go

start:
	./server

dev: 
	air
