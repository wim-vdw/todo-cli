FROM golang:1.22-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
#COPY internal ./internal
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-cli .

FROM alpine:3.20.3

ENV TIMEZONE=Europe/Brussels

RUN apk add alpine-conf

RUN setup-timezone -z $TIMEZONE

WORKDIR /

COPY --from=build-stage /todo-cli /todo-cli

ENTRYPOINT ["/todo-cli"]
