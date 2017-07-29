FROM golang:1.8
RUN mkdir -p /go/src/github.com/keang/goattache
WORKDIR /go/src/github.com/keang/goattache
COPY . .
RUN go-wrapper install
CMD ["go-wrapper", "run", "--port", "9292"]
