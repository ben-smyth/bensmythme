FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bensmythme /app/cmd/serve/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/bensmythme /app/bensmythme
COPY --from=builder /app/web /app/web
COPY --from=builder /app/spec.yaml /app/spec.yaml

RUN chmod +x /app/bensmythme

WORKDIR /app

CMD ["./bensmythme"]
