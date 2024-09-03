FROM golang:latest as builder

RUN mkdir /app
ADD . /app/
WORKDIR /app


RUN go env -w GOCACHE=/go-cache
RUN go env -w GOMODCACHE=/gomod-cache

# Сборка приложения с использованием кеша
RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache \


    go mod tidy

RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache \


    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o falconauth ./cmd/main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/tg_bot /app/tg_bot

EXPOSE 8052


CMD ["./tgbot"]

