set -ex

psql -U postgres -h localhost -p 5432 twitter \
    -c "CREATE PUBLICATION twitter_db_pub FOR ALL TABLES;"

pg_dump -U postgres -h localhost -s -p 5432 twitter | psql -U postgres -h localhost -p 5433 twitter_slave

psql -U postgres -h localhost -p 5433 twitter_slave \
    -c "CREATE SUBSCRIPTION twitter_db_sub CONNECTION 'dbname=twitter host=database user=postgres password=postgres port=5432' PUBLICATION twitter_db_pub;"
