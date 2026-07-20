#!/usr/bin/env bash

set -euo pipefail

echo "==> Starting BeaconLink development environment"

# -----------------------------------------------------------------------------
# Backend
# -----------------------------------------------------------------------------

echo "Starting API..."

go run ./cmd/api &
API_PID=$!

# -----------------------------------------------------------------------------
# Frontend
# -----------------------------------------------------------------------------

if [ -f "web/package.json" ]; then
    (
        cd web
        npm run dev
    ) &
    WEB_PID=$!
fi

cleanup() {
    echo
    echo "Stopping development environment..."

    kill "${API_PID}" 2>/dev/null || true

    if [ -n "${WEB_PID:-}" ]; then
        kill "${WEB_PID}" 2>/dev/null || true
    fi
}

trap cleanup EXIT INT TERM

wait