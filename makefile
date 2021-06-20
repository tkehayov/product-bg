docker-build:
	env GOOS=linux GOARCH=amd64 go build -o ../docker/volumes/merchants/app/ ./main.go