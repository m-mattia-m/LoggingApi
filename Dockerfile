FROM golang:latest AS builder

WORKDIR /build

COPY . /build

COPY /sites /build

RUN go build -o /build/main

FROM alpine:latest

RUN apk add --no-cache libc6-compat 

WORKDIR /app

ENV GIN_MODE=release

COPY --from=builder /build/main /app

CMD [ "/app/main" ]