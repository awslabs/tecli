FROM ubuntu:latest
WORKDIR /tmp
COPY dist/tecli-linux-amd64 tecli
RUN ./tecli