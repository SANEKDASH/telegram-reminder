FROM golang:1.26-alpine3.24@sha256:70b46548e42db77e0966aaf3619fd068734dc6c77584d526b91126504fd95816 AS deps

WORKDIR /telegram-bot
COPY go.mod go.sum ./
RUN go mod download


FROM golang:1.26-alpine3.24@sha256:70b46548e42db77e0966aaf3619fd068734dc6c77584d526b91126504fd95816 AS builder

WORKDIR /telegram-bot

COPY --from=deps /telegram-bot/go.mod /telegram-bot/go.sum ./
COPY --from=deps /go/pkg /go/pkg
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o telegram-bot-service


FROM alpine:3.23@sha256:fd791d74b68913cbb027c6546007b3f0d3bc45125f797758156952bc2d6daf40

RUN apk update && apk upgrade

RUN adduser -D -u 10001 -h /home/telegram-bot-user telegram-bot-user
WORKDIR /home/telegram-bot-user

COPY --from=builder /telegram-bot/telegram-bot-service ./
USER telegram-bot-user:telegram-bot-user

ENTRYPOINT ["/home/telegram-bot-user/telegram-bot-service"]
