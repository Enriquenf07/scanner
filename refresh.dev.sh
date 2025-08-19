#!/bin/bash
docker compose -f docker-compose.dev.yml down
docker rmi server
docker build --no-cache -t server:latest ./backend
docker stop frontend
docker rm frontend
docker rmi frontend --force
docker build --no-cache -f dockerfile.react -t frontend:latest ./

docker compose -f docker-compose.dev.yml up -d
