#!/usr/bin/env bash
set -euo pipefail

API_ADDR="${API_ADDR:-127.0.0.1:8080}" go run ./cmd/api
