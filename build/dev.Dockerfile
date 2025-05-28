FROM golang:1.24.3-alpine AS app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/air-verse/air@latest

COPY . .

ENV APP_ENV=local

CMD ["air", "-c", "./air.toml"]
