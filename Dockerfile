FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bensmythme

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/bensmythme /app/bensmythme
COPY --from=builder /app/web /app/web

RUN chmod +x /app/bensmythme

WORKDIR /app
EXPOSE 8080

CMD ["./bensmythme"]