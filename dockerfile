FROM        --platform=linux/amd64 golang:1.20
RUN         mkdir -p /app
WORKDIR     /app
COPY        . .

ENV USER=""
ENV USERPWD=""
ENV TZ="Asia/Taipei"

RUN         go mod download
RUN         go build -o app
ENTRYPOINT  ["./app"]

