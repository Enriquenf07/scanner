#!/bin/bash
docker compose down
docker rmi server
docker build -t server:latest ./backend
docker compose up -d
