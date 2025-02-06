FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOARCH=arm64 GOOS=linux go build -o application .

FROM alpine:3.21

ENV HOST_PROC=/host/proc
#ENV HOST_SYS=/host/sys
#ENV HOST_ETC=/host/etc
#ENV HOST_VAR=/host/var
#ENV HOST_RUN=/host/run
#ENV HOST_DEV=/host/dev
#ENV HOST_ROOT=/

WORKDIR /root/

COPY --from=builder /app/application .

EXPOSE 8080

CMD ["./application"]