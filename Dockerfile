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
ENV DATABASE_URL="postgresql://vsense:LKy4SHEP00O61gMj4uKQtI2S0L9iEOiz@dpg-cuc7aplumphs73ducp3g-a.oregon-postgres.render.com/telephone_3fr8"
ENV SECRETE_KEY="vsense"


ENTRYPOINT [ "./main" ]
