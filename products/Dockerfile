FROM golang:1.16

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
RUN go get go.mongodb.org/mongo-driver/bson
RUN go get go.mongodb.org/mongo-driver/mongo
CMD ["/app/main"]
