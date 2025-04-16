#!/bin/bash

export PGPASSWORD='pass123'
export ROOT_USER_PSQL='postgres'
export DB_PORT=5432
export DB_HOST=$(hostname -I | awk '{print $1}')
export DB_TYPE='postgres'
export DB_SSLMODE=disable

export LOGIN_PRODUCT_SERVICE='product_service'
export DB_PRODUCT_SERVICE='product_service'
export PASSWORD_PRODUCT_SERVICE='nahJ1iyeehei4ori6osh5I'
export JWT_SECRET_KEY='kahxein2Theey2Jae8Doh1'
