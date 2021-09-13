docker-build:
	env GOOS=linux GOARCH=amd64 go build -o ../docker/volumes/ftp-client/app/ ./main.go

run:
	env PORT=10091 go run main.go