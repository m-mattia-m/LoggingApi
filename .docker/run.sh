#!/usr/bin/env bash

# run ./.docker/run.sh

VERSION=$(git rev-parse --abbrev-ref HEAD)

docker run --rm -p 8080:8080 logging-api:$VERSION
