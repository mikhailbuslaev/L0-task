#!/bin/bash
echo "Running container...";
docker run --name orders_test -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres;
docker cp ./db.sql orders_test:/db.sql;
docker exec -it orders_test psql -U postgres -d orders_test  -f /db.sql;