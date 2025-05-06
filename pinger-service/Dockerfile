FROM golang:1.24-alpine AS build

WORKDIR /src

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o pinger-service

FROM alpine:edge

WORKDIR /src

COPY --from=build /src/pinger-service .

RUN chmod +x /src/pinger-service

RUN apk --no-cache add ca-certificates

EXPOSE 1234

ENTRYPOINT [ "/src/pinger-service" ]