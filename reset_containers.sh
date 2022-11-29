set -ex

docker container rm twc-postgres_database
docker volume rm twitter-clone-remake_postgresql_data
docker container ls -a
