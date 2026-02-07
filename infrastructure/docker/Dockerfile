# Build

FROM golang:1.25-alpine AS builder
ENV CGO_ENABLED=0
RUN apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY src ./src

RUN go build -ldflags="-s -w" -o seryn ./src

# Runtime
FROM alpine:3.23

RUN addgroup -S seryn && adduser -S seryn -G seryn
RUN apk add --no-cache git
WORKDIR /app
COPY --from=builder /app/seryn /usr/local/bin/seryn

USER seryn
ENTRYPOINT ["seryn"]