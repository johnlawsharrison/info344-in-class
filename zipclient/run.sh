#!/usr/bin/env bash
set -e
./build.sh
docker run -p 4200:4200 -d johnlawsharrison/zipclient
