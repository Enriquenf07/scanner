#!/bin/bash
docker compose down
# docker rmi server
# docker build --no-cache -t server:latest ./backend

docker compose up -d
