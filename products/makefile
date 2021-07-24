docker-build:
	env GOOS=linux GOARCH=amd64 go build -o ../docker/volumes/products/app/ ./main.go;
	env GOOS=linux GOARCH=amd64 go build -o ../docker/volumes/products/app/ ./product-provider.go;

run:
	go run main.go & go run products-provider.go