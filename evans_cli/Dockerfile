FROM ubuntu:20.04

WORKDIR /usr/src
COPY ./common/protos ./protos

WORKDIR /usr/src/app

RUN apt update && apt -y install curl && apt -y install dnsutils
RUN curl -L 'https://github.com/ktr0731/evans/releases/download/v0.10.6/evans_linux_amd64.tar.gz' | tar xvzf -
RUN mv evans /usr/local/bin