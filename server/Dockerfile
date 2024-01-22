FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk --no-cache add bash git make gcc gettext musl-dev

# dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# build
COPY . .
RUN go build -o ./bin/main cmd/main.go

FROM alpine

WORKDIR /app
 
COPY --from=builder /app/bin/main /

COPY --from=builder /app/.env .env

CMD ["/main"]