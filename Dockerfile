FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/main

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/bin/main .

ENV SERVER_ADDRESS="0.0.0.0:8080"
ENV DATABASE_URL="postgresql://vsense:GmTaf9S8uVo6FERprSKBZpDZNpYLtkJG@dpg-csqo6i5umphs73d5fllg-a.singapore-postgres.render.com/telephone"
ENV SECRETE_KEY="vsense"


ENTRYPOINT [ "./main" ]
