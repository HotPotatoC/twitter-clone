#!/usr/bin/env bash

set -euo pipefail

function run_force_migrations() {
    local db_url="$1"
    local version="$2"

    migrate -database "$db_url" -path database/migrations force "$version"
}

if [ -z "${1-}" ]; then
    echo "Missing database url argument"
    exit 1
fi

if [ -z "${2-}" ]; then
    echo "Missing migration version"
    exit 1
fi

run_force_migrations $1 $2
exit 0
