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
ENV DATABASE_URL="postgresql://vsync:gDxtodwrh41K2Mtkps5rowGdLe0ii1us@dpg-ctn3vja3esus739tv4p0-a.oregon-postgres.render.com/telephone_5fwr"
ENV SECRETE_KEY="vsense"


ENTRYPOINT [ "./main" ]
