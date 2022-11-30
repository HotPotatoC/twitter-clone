set -ex

psql -U postgres -h localhost -p 5432 \
    -c "CREATE PUBLICATION twitter_db_pub FOR ALL TABLES;"

pg_dump -U postgres -h localhost -s public -s -p 5432 twitter | psql -U postgres -h localhost -p 5433 twitter_slave

psql -U postgres -h localhost -p 5433 \
    -c "CREATE SUBSCRIPTION twitter_db_sub connection 'dbname=twitter host=host.docker.internal user=postgres password=postgres port=5432' publication twitter_db_pub;"
