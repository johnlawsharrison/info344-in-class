#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t johnlawsharrison/zipsvr .
docker push johnlawsharrison/zipsvr
go clean