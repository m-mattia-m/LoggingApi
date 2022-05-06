#!/usr/bin/env bash

# run ./.docker/build.sh

VERSION=$(git rev-parse --abbrev-ref HEAD)

docker build . -t logging-api:$VERSION

