FROM golang
ENV GOPATH="/go"
RUN ["mkdir", "-p", "/go/src/github.com/openipd/rio-demo"]
ADD * /go/src/github.com/openipd/rio-demo/
WORKDIR /go/src/github.com/openipd/rio-demo
RUN ["go", "build", "-o", "demo"]
CMD ["./demo"]
