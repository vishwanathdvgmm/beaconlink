#!/usr/bin/env bash

set -euo pipefail

echo "==> Building BeaconLink"

mkdir -p bin

# -----------------------------------------------------------------------------
# Backend
# -----------------------------------------------------------------------------

go build -o bin/beacon ./cmd/beacon

go build -o bin/agent ./cmd/agent

go build -o bin/relay ./cmd/relay

go build -o bin/api ./cmd/api

# -----------------------------------------------------------------------------
# Frontend
# -----------------------------------------------------------------------------

if [ -f "web/package.json" ]; then
    (
        cd web
        npm run build
    )
fi

echo
echo "Build complete."