FROM golang:alpine3.10 as build

ENV GO111MODULE=on

WORKDIR /go/src/hello-kubecon

COPY main.go .

RUN go mod init && \
    go build -o /go/bin/kubecon main.go && \
    ls /go/bin/

FROM alpine:3.10.3

COPY --from=build /go/bin/kubecon /

CMD ["/kubecon"]

EXPOSE 8080