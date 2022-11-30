set -ex

sudo rm -rf master slave

docker container rm twc-postgres_database twc-postgres_database-slave
docker volume rm twitter-clone-remake_postgres-data-master twitter-clone-remake_postgres-data-slave

docker container ls -a
