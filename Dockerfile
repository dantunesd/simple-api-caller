FROM golang:1.18-alpine as build

WORKDIR /app

COPY . .

RUN go build

FROM alpine:latest as app

WORKDIR /bin

COPY --from=build /app/simple-api-caller simple-api-caller

CMD ["./simple-api-caller"]