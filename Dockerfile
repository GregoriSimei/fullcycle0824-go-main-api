FROM golang:1.23

RUN apt update && apt install -y openssl procps

WORKDIR /usr/src/app

CMD tail -f /dev/null
