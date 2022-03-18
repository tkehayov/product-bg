docker-build:
	env GOOS=linux GOARCH=amd64 go build -o ../../docker/volumes/merchat-catalog/app/ ../cmd/main.go

run:
	env PORT=9091 go run ../cmd/main.go