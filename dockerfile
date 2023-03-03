FROM        --platform=linux/amd64 golang:1.20
RUN         mkdir -p /app
WORKDIR     /app
COPY        . .

ENV USERNAME=""
ENV USERPWD=""

RUN         go mod download
RUN         go build -o app
ENTRYPOINT  ["./app"]

