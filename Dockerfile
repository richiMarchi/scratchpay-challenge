FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV CGO_ENABLED 0
RUN go test ./...

RUN go build -o /users-api ./cmd/restserver

#---------------------------------------
FROM alpine:latest

WORKDIR /

COPY --from=build /users-api /users-api

RUN mkdir /misc && addgroup -S appgroup && adduser -S appuser -G appgroup && chown -R appuser /misc
USER appuser

EXPOSE 8080

ENTRYPOINT ["/users-api"]
