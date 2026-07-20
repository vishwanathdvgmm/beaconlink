#!/usr/bin/env bash

set -euo pipefail

echo "==> Running code generation"

# -----------------------------------------------------------------------------
# Go Generate
# -----------------------------------------------------------------------------

go generate ./...

# -----------------------------------------------------------------------------
# sqlc
# -----------------------------------------------------------------------------

if command -v sqlc >/dev/null 2>&1; then
    echo "==> sqlc generate"
    sqlc generate
fi

echo
echo "Generation complete."