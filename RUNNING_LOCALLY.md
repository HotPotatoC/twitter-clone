# Installation - Running locally 💻

Here you will understand how to run and setup the development environment for twitterclone

> NOTE: To run this app locally you need to have an AWS S3 Bucket available so that uploading images will work. (Might add Cloudinary and on-disk implementation in the future)

## 1. Using docker-compose

***Prerequisites***
- [Docker](https://docker.com/)
- [AWS Access Key](https://docs.aws.amazon.com/powershell/latest/userguide/pstools-appendix-sign-up.html) and [S3 Bucket](https://aws.amazon.com/s3/)

Create .env file in `configs/` directory by copying `configs/.env.example` and setup the environment variables

- Running the containers with `docker-compose up -d`
- To stop the containers `docker-compose stop`

## 2. Manually

***Prerequisites***
- [Golang](golang.org)
- [PostgreSQL](postgresql.org)
- [Redis](redis.io)
- [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [AWS Access Key](https://docs.aws.amazon.com/powershell/latest/userguide/pstools-appendix-sign-up.html) and [S3 Bucket](https://aws.amazon.com/s3/)
- [air (optional hot-reload)](https://github.com/cosmtrek/air)

###  Local Rest API Server Development

#### PostgreSQL
Install PostgreSQL:
- **macOS**: Run `brew install postgresql`.
- **Windows**: Follow [this](https://www.postgresqltutorial.com/install-postgresql/) guide.
- **Linux**: Follow [this](https://www.postgresqltutorial.com/install-postgresql-linux/) guide.


Create a database named `twitterclone`

```sh
$ psql

postgres=# CREATE DATABASE twitterclone;
```

Install golang-migrate [here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

Run the migrations.

```sh
$ ./scripts/run_migrations.sh Your_PostgreSQL_URL

# Example PostgreSQL url: postgres://postgres:postgres@127.0.0.1:5432/twitterclone\?sslmode=disable
```

#### Redis
Install Redis:
- **macOS**: Run `brew install redis`.
- **Windows**: Follow [this](https://redis.io/download#installation) guide.
- **Linux**: Follow [this](https://redis.io/download#installation) guide.

#### Golang
Install Golang:
- **macOS**: Run `brew install golang`.
- **Windows**: Follow [this](https://golang.org/dl/) guide.
- **Linux**: Follow [this](https://golang.org/dl/) guide.


Navigate to `/configs` and set the following environment variables:

```
APP_NAME=twitter-clone
APP_HOST=127.0.0.1
APP_PORT=5000
APP_DOMAIN=localhost
DEBUG=false

ACCESS_TOKEN_SECRET=
REFRESH_TOKEN_SECRET=
# Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"
# 15m -> 15 minutes
ACCESS_TOKEN_DURATION=15m
# 168h -> 7 days
REFRESH_TOKEN_DURATION=168h

DEFAULT_AVATAR_URL=

# 2.5 mb in bytes
MAX_UPLOAD_SIZE=2621440

# For attaching images in tweets
# 32 mb in bytes
MAX_TWEET_ATTACHMENT_SIZE=33554432

DB_USER=postgres
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=5432
DB_DATABASE=twitterclone

REDIS_ADDR=127.0.0.1:6379
REDIS_PASSWORD=
REDIS_KEY_DELIMITER=::

AWS_REGION=
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
AWS_S3_BUCKET_NAME=

# [Optional] can leave empty here
AWS_SESSION_TOKEN=
```

**Running the server**

```sh
$ go run cmd/rest/main.go --prefork

# Using hot-reloading (Optional must have air installed)
$ make dev
```

### Local Frontend Development

Navigate to `/web`

- Run `yarn` to install the dependencies
- Run `yarn dev` to run the frontend
