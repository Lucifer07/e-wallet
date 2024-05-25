FROM golang:1.18-alpine AS builder



WORKDIR /app



# Download Go modules

COPY go.mod go.sum /app/

RUN go mod download



COPY . /app/



RUN CGO_ENABLED=0 GOOS=linux go build -o library-api .





FROM alpine:3 AS dev



WORKDIR /app



COPY --from=builder /app/library-api /app/.env /app/



EXPOSE 8080



ENTRYPOINT [ "./library-api" ]