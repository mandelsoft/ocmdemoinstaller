FROM golang:1.17 AS builder

WORKDIR /go/src/github.com/mandelsoft/ocmdemoinstaller/
COPY go .
#COPY go/pkg pkg
RUN go get -d ./...
RUN go build -o /main ./cmd

FROM alpine

COPY content ocm
COPY --from=builder /main /ocm/run
ENTRYPOINT [ "/ocm/run" ]


