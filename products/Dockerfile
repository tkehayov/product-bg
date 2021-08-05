FROM golang:1.16-alpine3.14

RUN mkdir /app
#COPY go.mod /app
ADD . /app
#WORKDIR /app
#RUN go mod tidy
#RUN go mod download
RUN go build -o main .

CMD ["/app/main"]
