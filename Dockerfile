FROM ubuntu:latest AS ubuntu
WORKDIR /tmp
COPY build/tfe-cli .
RUN ./tfe-cli

FROM ubuntu:bionic AS bionic
WORKDIR /tmp
COPY build/tfe-cli .
RUN ./tfe-cli