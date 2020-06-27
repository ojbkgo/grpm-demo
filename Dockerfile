FROM golang:1.11.13

WORKDIR $GOPATH/src/github.com/ojbkgo/grpc-demo
ADD . $GOPATH/src/github.com/ojbkgo/grpc-demo
ENV GO111MODULE=off
RUN go build -o bin/demo main.go \
    && go build -o bin/demo-cli cmd/demo.go

CMD ["bin/demo"]