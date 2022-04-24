build: 
	go build -o server main.go

run: build
	./server

watch:
	ulimit -n 10000
	reflex -s -r '\.go$$' make run