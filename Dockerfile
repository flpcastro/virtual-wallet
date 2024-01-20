FROM golang:1.18.2-alpine3.16 as base

RUN apk update

WORKDIR /src/virtualwallet

ADD . .

RUN go mod download

RUN go build -o virtualwallet ./cmd/api

FROM alpine:3.16 as binary

WORKDIR /src/app

COPY --from=base /src/virtualwallet/virtualwallet .

EXPOSE 3000

CMD ["/src/app/virtualwallet"]