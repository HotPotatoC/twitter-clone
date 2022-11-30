set -ex

pg_dump -U postgres -h localhost -s public -s twitter -p 5432 | psql -U postgres -h localhost -p 5433 twitter_slave
