#!/usr/bin/env bash

set -euo pipefail

private_key_out="configs/server.key"
cert_key_out="configs/server.crt"

echo "==> Generating RSA private key"
openssl genrsa -out $private_key_out 2048

echo "==> Generating Cert key"
openssl req -new -x509 -sha256 -key $private_key_out -out $cert_key_out -days 3650

exit 0
