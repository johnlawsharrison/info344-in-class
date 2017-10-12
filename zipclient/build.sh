#!/usr/bin/env bash
set -e
docker build -t johnlawsharrison/zipclient .
docker push johnlawsharrison/zipclient