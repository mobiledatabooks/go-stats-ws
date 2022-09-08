#!/usr/bin/env bash

# Pushing Docker Image to Container Registry
echo docker build -t gcr.io/${PROJECT}/${SERVICE_NAME} .

docker build --platform linux/amd64 -t gcr.io/${PROJECT}/${SERVICE_NAME} .
docker push gcr.io/${PROJECT}/${SERVICE_NAME}
# 

# gcloud run deploy ${SERVICE_NAME} --image=gcr.io/${PROJECT}/${SERVICE_NAME} --platform managed --project ${PROJECT} --allow-unauthenticated --memory 512Mi --port=8085
