set -ex

docker-compose down || true
docker volume rm twitter-clone-remake_postgres-data-master twitter-clone-remake_postgres-data-slave || true

docker container ls -a
