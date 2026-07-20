#!/usr/bin/env bash

set -euo pipefail

echo "==> Bootstrapping BeaconLink development environment"

# -----------------------------------------------------------------------------
# Verify required tools
# -----------------------------------------------------------------------------

required_tools=(
  go
  git
)

for tool in "${required_tools[@]}"; do
  if ! command -v "$tool" >/dev/null 2>&1; then
    echo "Missing required tool: $tool"
    exit 1
  fi
done

# -----------------------------------------------------------------------------
# Go Modules
# -----------------------------------------------------------------------------

echo "==> Downloading Go modules"

go mod download

# -----------------------------------------------------------------------------
# Frontend
# -----------------------------------------------------------------------------

if [ -f "web/package.json" ]; then
    echo "==> Installing frontend dependencies"

    (
        cd web
        npm install
    )
fi

echo
echo "Bootstrap complete."