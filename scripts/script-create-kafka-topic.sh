#!/bin/bash

docker exec -it kafka-1 bash -c \
    "kafka-topics --create --bootstrap-server kafka-1:29091 --replication-factor 3 --partitions 3 --topic payment_events"
