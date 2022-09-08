#!/usr/bin/env bash
docker stop mobiledatabooks-stats-service
docker rm mobiledatabooks-stats-service
docker build -f Dockerfile -t mobiledatabooks-stats-service .

docker run -d \
    -p 8081:8085 \
    --name mobiledatabooks-stats-service \
    mobiledatabooks-stats-service
