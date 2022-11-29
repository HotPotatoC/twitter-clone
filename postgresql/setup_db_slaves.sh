set -ex

# postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/

psql -U postgres -h localhost -p 5432 -c "CREATE USER replicator WITH REPLICATION ENCRYPTED PASSWORD '$(
    tr </dev/urandom -dc _A-Z-a-z-0-9 | head -c${1:-32}
    echo
)';"

psql -U postgres -h localhost -p 5432 -c "SELECT * FROM pg_create_physical_replication_slot('replication_slot_slave1');"

docker exec twc-postgres_database pg_basebackup -D /tmp/postgresslave -S replication_slot_slave1 -X stream -P -U replicator -Fp -R -h localhost -p 5432
