# syntax=docker/dockerfile:1

FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o worker ./cmd/worker/main.go

FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=builder /app/api ./api
COPY --from=builder /app/worker ./worker
CMD ["/app/api"] 