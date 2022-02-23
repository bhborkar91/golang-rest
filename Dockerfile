FROM golang:1.13-alpine as base

WORKDIR /work

COPY local_ca_files/* /usr/local/share/ca-certificates/
RUN update-ca-certificates

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./
RUN go build -o /app
RUN chmod +x /app

FROM alpine:3.12
# ENV GIN_MODE=release
COPY --from=base /app /app
EXPOSE 4000
ENTRYPOINT ["/app"]