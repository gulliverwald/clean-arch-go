FROM golang:1.18-alpine

WORKDIR /app

COPY . .
RUN go mod tidy

CMD ["go", "run", "app/main.go"]
