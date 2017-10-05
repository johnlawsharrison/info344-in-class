#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t johnlawsharrison/testserver .
docker push johnlawsharrison/testserver