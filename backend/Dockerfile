FROM golang:1.24.10-alpine as builder
# Todo task [if required will be adding]
# 1 install dependencies
# 2 setup working directory
# 3 copy go mod files
# 4 download dependencies
# 5 copy source code
# 6 build application
# 7 final stage 
# 8 install runtime dependencies
# 9 create non-root users
# 10  set up working directory
# 11  copy binary from builder
# 12 Create directories or storage
# 13 switch to non-root user
# 14 Expose port
# 15 health check

RUN apk add --no-cache git make gcc musl-dev
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

WORKDIR /home/appuser

COPY --from=builder /app/server .
RUN mkdir - p storage/pdf storage/thumbnails && \
    chown -R appuser:appuser storage

USER appuser

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/api/health || exit 1

CMD ["./server"]
