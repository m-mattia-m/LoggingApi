FROM golang:latest AS builder

WORKDIR /build

COPY . /build

RUN go build -o /build/main

FROM alpine:latest

RUN apk add --no-cache libc6-compat 

WORKDIR /app

ENV GIN_MODE=release

COPY --from=builder /build/main /app
COPY --from=builder /build/sites /app/sites

EXPOSE 8080

CMD [ "/app/main" ]
