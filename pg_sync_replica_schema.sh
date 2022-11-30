set -ex

pg_dump -U postgres -h localhost -s -p 5432 twitter | psql -U postgres -h localhost -p 5433 twitter_slave
