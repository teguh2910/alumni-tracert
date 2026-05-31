# Stage 1: Build
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/alumni server.go

# Stage 2: Run
FROM alpine

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/alumni /app/
COPY stempel-sign.png /app/

EXPOSE 9099
EXPOSE 8099

CMD ["/app/alumni"]
