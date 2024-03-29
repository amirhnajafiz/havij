FROM golang:1.17-alpine as builder

WORKDIR /app/src/

COPY . .

WORKDIR /app/src

RUN go build -o ./runner

FROM alpine:latest

WORKDIR /app/

COPY --from=builder ./runner .

ENTRYPOINT ["./runner"]
