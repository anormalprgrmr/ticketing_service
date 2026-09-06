# Build stage
FROM golang:1.27.1-alpine3.23 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/main.go

FROM alpine:latest
WORKDIR /
COPY --from=build /app/app /
CMD ["./app"]
