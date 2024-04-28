FROM golang:1.22-alpine3.19 AS builder
WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /go/bin/snorlax 

FROM alpine:3.19
COPY --chown=65534:65534 --from=builder /go/bin/snorlax .
USER 65534

ENTRYPOINT [ "./snorlax" ]