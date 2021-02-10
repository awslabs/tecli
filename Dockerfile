FROM ubuntu:18.04

RUN apt-get -qq update -y
RUN apt-get -qq install ca-certificates
ARG TFC_TEAM_TOKEN
WORKDIR /opt/
COPY dist/tecli-linux-amd64 ./tecli

RUN chmod +x tecli* \
    && ./tecli configure create --profile=default --enabled=true --team-token=${TFC_TEAM_TOKEN} --description="org/foo-org" \
    && ./tecli workspace list --organization=foo-org

RUN ls -ltha /etc

