#!/usr/bin/env bash

set -euxo pipefail

export WORKSPACE=$(pwd)
export DEBIAN_FRONTEND=noninteractive

sudo apt-get update
sudo apt-get -y install --no-install-recommends \
	bash-completion \
    python3-pip \
    make

pip3 install --user -r "${WORKSPACE}/.devcontainer/requirements.txt"

export GOPROXY=direct
go install golang.org/x/tools/cmd/goimports@latest

clear && devcontainer-info
