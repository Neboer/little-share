FROM golang:stretch
WORKDIR /go/src/app
VOLUME /go/src/app/files
COPY . .
ENV GIN_MODE=release
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io
RUN go get -d -v ./...
RUN go install -v ./...
CMD go run main.go