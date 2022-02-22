FROM golang:1.13-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

RUN go build -o /app

FROM scratch
EXPOSE 4000
ENTRYPOINT ["/app"]