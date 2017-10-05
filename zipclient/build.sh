#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t johnlawsharrison/zipclient .
docker push johnlawsharrison/zipclient
go clean