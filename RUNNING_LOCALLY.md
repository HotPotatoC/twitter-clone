# Running Locally 💻

Here you will understand how to run and setup the development environment for twitterclone in your local machine.

> This guide is intended for Linux or MacOS users. For Windows users you could use something like [WSL](https://learn.microsoft.com/en-us/windows/wsl/install).

***Prerequisites***
- [Docker](https://docker.com/)

## 1. PostgreSQL Setup

Start the containers

```bash
make start
# or
docker-compose up
```

Run the replication setup script

```bash
make db-setup-logical-replication
# or
chmod +x ./postgresql/setup_logical_replication.sh
./postgresql/setup_logical_replication.sh
```

You're done!

If you have any problems feel free to open an [issue](https://github.com/HotPotatoC/twitter-clone/issues/new).