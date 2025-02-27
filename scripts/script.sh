#!/bin/bash


if [[ "$1" == "--delete" ]]; then
    echo "Удаление базы данных..."
    source /home/zangar/Documents/product_service/scripts/script-delete-psql.sh
else
    echo "Создание и запуск базы данных..."
    /home/zangar/Documents/product_service/scripts/script-psql-run.sh
    sleep 4
    /home/zangar/Documents/product_service/scripts/script-create-db.sh
fi
