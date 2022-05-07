start-app:
	go run main.go

start-worker:
	go run queue-worker.go

build-app:
	go build main.go

build-worker:
	go build queue-worker.go
