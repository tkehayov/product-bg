docker-build:
	env GOOS=linux GOARCH=amd64 go build -o ../../docker/volumes/products-collector/app/ ../cmd/main.go;

run:
	env MONGO_URL=mongodb://localhost:37017/products\
	env PORT=9092 go run ../cmd/main.go