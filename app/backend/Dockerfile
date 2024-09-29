FROM golang:1.23 AS builder
WORKDIR /app/build

COPY ./src/ /app/build/.
RUN go install
RUN go build -o bin

FROM debian:stable-slim AS runtime
WORKDIR /app

COPY --from=builder /app/build/bin ./bin

ENTRYPOINT ["/app/bin"]
