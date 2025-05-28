FROM golang:1.24.3-alpine AS app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV APP_ENV=local

CMD ["go", "run", "./main.go"]
