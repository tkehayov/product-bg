FROM golang:1.16

RUN mkdir /app
ADD products /app
ADD merchants /app
WORKDIR /app
RUN go build -o products .
RUN go build -o merchants .
RUN go get go.mongodb.org/mongo-driver/bson
RUN go get go.mongodb.org/mongo-driver/mongo
CMD ["/app/products","/app/merchants"]
