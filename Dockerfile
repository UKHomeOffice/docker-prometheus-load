FROM golang:1.14 as builder
ADD . /src
WORKDIR /src
RUN CGO_ENABLED=0 go build -o prometheus-load .

FROM alpine:3.12
WORKDIR /app
COPY --from=builder /src/prometheus-load .
ENTRYPOINT ["/app/prometheus-load"]