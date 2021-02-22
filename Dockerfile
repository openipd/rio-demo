FROM golang:1.14
ENV GOPATH="/go"
RUN ["mkdir", "-p", "/go/src/github.com/openipd/rio-demo"]
COPY * /go/src/github.com/openipd/rio-demo/
WORKDIR /go/src/github.com/openipd/rio-demo
RUN ["go", "build", "-o", "demo"]
CMD ["./demo"]
