ARG BASE_IMG_TAG=nginx:latest
FROM ${BASE_IMG_TAG}

COPY nginx.conf /etc/nginx/nginx.conf