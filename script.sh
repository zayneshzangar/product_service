#!/bin/bash

./env.sh
./script-psql-run.sh
sleep 2
./script-create-db.sh
