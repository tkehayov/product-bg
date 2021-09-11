docker-build:
	env GOOS=linux GOARCH=amd64 go build -o ../../docker/volumes/merchants/app/ ../main.go
	env GOOS=linux GOARCH=amd64 go build -o ../../docker/volumes/merchants-consumer/app/ ../merchant-consumer.go;

run:
	env MONGO_URL=mongodb://localhost:37018/merchants\
	env CONSUMER_PORT=50051\
	env PORT=9091 go run main.go