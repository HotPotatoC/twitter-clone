#!/usr/bin/env bash

set -euo pipefail

function run_migrations() {
    local db_url="$1"
    local method="$2"

    migrate -database "$db_url" -path database/migrations "$method"
}

if [ -z "${1-}" ]; then
    echo "Missing database url argument"
    exit 1
fi

run_migrations $1 ${2:-'up'}
exit 0
