FROM golang:1.20.5-alpine3.18
ENV ROOT=/app
WORKDIR ${ROOT}
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o main

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /app/main /app/.env ./
EXPOSE 8080
CMD ["./main"]