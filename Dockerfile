FROM golang:latest AS builder

WORKDIR /build

COPY . /build

RUN go mod download

RUN go build .

EXPOSE 8080

FROM alpine:latest

RUN apk add --no-cache libc6-compat 

WORKDIR /app

COPY --from=builder /build/main /app/

EXPOSE 8080

CMD [ "/app/main" ]

