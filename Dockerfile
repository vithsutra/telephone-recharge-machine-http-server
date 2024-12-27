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
ENV DATABASE_URL="postgresql://vsync:q7u1I6gsoPHT1UyViPAgJ4fFeatjyRlF@dpg-ctn3npi3esus739tte40-a.oregon-postgres.render.com/telephone_hqtb"
ENV SECRETE_KEY="vsense"


ENTRYPOINT [ "./main" ]
