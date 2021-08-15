FROM ubuntu:20.04

WORKDIR /usr/src
COPY ./common/protos ./protos

RUN apt update && apt -y install curl
RUN curl -L 'https://github.com/ktr0731/evans/releases/download/0.10.0/evans_linux_amd64.tar.gz' | tar xvzf -
RUN mv evans /usr/local/bin