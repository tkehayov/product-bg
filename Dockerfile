FROM golang:1.16

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]

#FROM golang:1.16
#WORKDIR /go/src/app
##RUN  go build -o ./docker_build/main ./cmd/main.go -w GO111MODULE=off build
#COPY docker_build/app /go/src/app

#FROM golang:1.16
#
#
#COPY . .
#
#RUN go get -d -v ./...
#RUN go install -v ./...

#CMD ["/go/src/app"]
#GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
#DOCKER_BUILD=$(shell pwd)/.docker_build
#DOCKER_CMD=$(DOCKER_BUILD)/product-bg
#$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) .

#FROM heroku/heroku:18-build as build
#
#COPY . /app
#WORKDIR /app
#
## Setup buildpack
#RUN mkdir -p /tmp/buildpack/heroku/go /tmp/build_cache /tmp/env
#RUN curl https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/go.tgz | tar xz -C /tmp/buildpack/heroku/go
#
##Execute Buildpack
#RUN STACK=heroku-18 /tmp/buildpack/heroku/go/bin/compile /app /tmp/build_cache /tmp/env
#
## Prepare final, minimal image
#FROM heroku/heroku:18
#
#COPY --from=build /app /app
#ENV HOME /app
#WORKDIR /app
#RUN useradd -m heroku
#USER heroku
#CMD /app/bin/product-bg~