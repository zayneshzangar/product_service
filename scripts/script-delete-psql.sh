#!/bin/bash

docker compose -f /home/zangar/Documents/product_service/scripts/docker-compose.yaml down
docker volume rm  scripts_psql-data # $(docker volume ls -q | grep scripts_psql-data)
