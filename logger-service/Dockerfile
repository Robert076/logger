FROM golang:1.24-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o logger-service

FROM alpine:edge

WORKDIR /src

COPY --from=build /src/logger-service .

RUN chmod +x /src/logger-service

RUN apk --no-cache add ca-certificates

EXPOSE 1234

ENTRYPOINT [ "/src/logger-service" ]