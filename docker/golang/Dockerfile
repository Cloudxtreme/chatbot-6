FROM golang:alpine

ENV GOPATH /go

RUN apk update -U \
    && apk add git

CMD ["tail", "-f", "/dev/stderr"]
