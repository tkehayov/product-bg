docker-build:
	env GOOS=linux GOARCH=amd64 go build -o ../../docker/volumes/products/app/ ../cmd/main.go;
	env GOOS=linux GOARCH=amd64 go build -o ../../docker/volumes/products-consumer/app/ ../cmd/products-consumer.go;

run:
	env MONGO_URL=mongodb://localhost:37017/products\
	env CONSUMER_PORT=50051\
	env PORT=9092 go run ../cmd/main.go & go run ../cmd/products-consumer.go