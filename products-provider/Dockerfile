FROM golang:1.16

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
RUN go mod tidy
CMD ["/app/main"]
