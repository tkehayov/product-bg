docker-build:
	env GOOS=linux GOARCH=amd64 go build -o ../docker/volumes/products/app/ ./main.go;
	env GOOS=linux GOARCH=amd64 go build -o ../docker/volumes/products-consumer/app/ ./products-consumer.go;

run:
	env MONGO_URL=mongodb://localhost:37017/merchants\
	env CONSUMER_PORT=50051\
	env PORT=9091 go run main.go & go run products-consumer.go