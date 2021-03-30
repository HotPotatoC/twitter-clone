#!/usr/bin/env bash

set -euo pipefail

function create_migration() {
    local name="$1"

    migrate create -ext sql -dir database/migrations -seq $name
}

if [ -z "${1-}" ]; then
    echo "Missing migration name argument"
    exit 1
fi

create_migration $1
exit 0
