package build

func DockerFileTem() string {
    t := `
FROM alpine:3.12

RUN apk add --no-cache tzdata

ENV TZ="Asia/Shanghai"

ARG HttpPort

RUN mkdir -p /app

WORKDIR /app
EXPOSE $HttpPort

COPY ./config /app/config
COPY ./bin/server /app/bin/server

#最终运行docker的命令
ENTRYPOINT  ["./bin/server"]

    `
    return t
}