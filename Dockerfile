# syntax=docker/dockerfile:1

FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./backend/cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o worker ./backend/cmd/worker/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o confirmation_worker ./backend/cmd/confirmation_worker/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o event_indexer ./backend/cmd/event_indexer/main.go

FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=builder /app/api ./api
COPY --from=builder /app/worker ./worker
COPY --from=builder /app/confirmation_worker ./confirmation_worker
COPY --from=builder /app/event_indexer ./event_indexer
CMD ["/app/api"] 