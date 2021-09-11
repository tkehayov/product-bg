docker-build:
	env GOOS=linux GOARCH=amd64 go build -o ../../docker/volumes/products-provider/app/ ../cmd/main.go

run:
	env PORT=9090 go run main.go
