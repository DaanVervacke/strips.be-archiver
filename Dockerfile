FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o strips.be-archiver cmd/main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/strips.be-archiver .
