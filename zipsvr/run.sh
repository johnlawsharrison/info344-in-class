#!/usr/bin/env bash
set -e
./build.sh
docker run -p 80:80 -d johnlawsharrison/zipsvr
