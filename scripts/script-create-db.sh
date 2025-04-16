#!/bin/bash

source /home/zangar/Documents/product_service/scripts/env.sh

psql -U $ROOT_USER_PSQL -p $DB_PORT -h $DB_HOST \
    -c "CREATE ROLE $DB_PRODUCT_SERVICE WITH PASSWORD '$PASSWORD_PRODUCT_SERVICE';"
psql -U $ROOT_USER_PSQL -p $DB_PORT -h $DB_HOST \
    -c "CREATE DATABASE $DB_PRODUCT_SERVICE;"
psql -U $ROOT_USER_PSQL -p $DB_PORT -h $DB_HOST \
    -c "ALTER ROLE $DB_PRODUCT_SERVICE WITH LOGIN;"
psql -U $ROOT_USER_PSQL -p $DB_PORT -h $DB_HOST \
    -c "GRANT ALL PRIVILEGES ON DATABASE $DB_PRODUCT_SERVICE TO $DB_PRODUCT_SERVICE;"
psql -U $ROOT_USER_PSQL -p $DB_PORT -h $DB_HOST \
    -c "GRANT pg_write_server_files TO $DB_PRODUCT_SERVICE;"
psql -U $ROOT_USER_PSQL -p $DB_PORT -h $DB_HOST \
    -c "GRANT pg_read_server_files TO $DB_PRODUCT_SERVICE;"
psql -U $ROOT_USER_PSQL -p $DB_PORT -h $DB_HOST \
    -c "GRANT CREATE ON SCHEMA public TO $DB_PRODUCT_SERVICE;"
psql -U $ROOT_USER_PSQL -p $DB_PORT -h $DB_HOST \
    -c "ALTER USER $DB_PRODUCT_SERVICE WITH SUPERUSER;"
PGPASSWORD=$PASSWORD_PRODUCT_SERVICE psql -U $DB_PRODUCT_SERVICE -p $DB_PORT -h $DB_HOST \
    -c 'CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    stock INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);'

# Вставка тестовых данных
PGPASSWORD=$PASSWORD_PRODUCT_SERVICE psql -U $DB_PRODUCT_SERVICE -p $DB_PORT -h $DB_HOST -d $DB_PRODUCT_SERVICE \
    -c "INSERT INTO products (name, description, price, stock) VALUES
    ('Ноутбук', 'Игровой ноутбук с RTX 4060', 12000.00, 10),
    ('Игровая клавиатура', 'Механическая клавиатура с RGB-подсветкой', 5000.00, 15),
    ('Беспроводная мышь', 'Игровая мышь с высокой точностью', 3500.00, 20),
    ('Монитор 27 дюймов', 'IPS-монитор с частотой 165 Гц', 30000.00, 8),
    ('Игровое кресло', 'Эргономичное кресло для геймеров', 15000.00, 5),
    ('Веб-камера 4K', 'Камера для стримов с автофокусом', 12000.00, 12);"
