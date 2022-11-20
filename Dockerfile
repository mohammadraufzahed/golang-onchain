FROM golang:1.19-bullseye as builder

WORKDIR /build

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN swag init --parseDependency --parseInternal --parseDepth 1 -g ./web/start.go

RUN go build -ldflags="-s -w" -o glassnode-api main.go

FROM debian:bullseye

WORKDIR /app

RUN apt update && apt install ca-certificates -y

ENV ENV=DOCKER
ENV HTTP_PROXY=http://glassnode:HBkYexQ2WspQUv5@82.115.16.18:8888
ENV HTTPS_PROXY=http://glassnode:HBkYexQ2WspQUv5@82.115.16.18:8888

COPY --from=builder /build/glassnode-api .
COPY --from=builder /build/docker.env .

CMD [ "/app/glassnode-api" ]

EXPOSE 3000
