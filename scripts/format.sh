#!/usr/bin/env bash

set -euo pipefail

echo "==> Formatting source code"

# -----------------------------------------------------------------------------
# Go Formatting
# -----------------------------------------------------------------------------

gofmt -w .

if command -v goimports >/dev/null 2>&1; then
    goimports -w .
fi

# -----------------------------------------------------------------------------
# Frontend Formatting
# -----------------------------------------------------------------------------

if [ -f "web/package.json" ]; then
    (
        cd web
        npm run format
    )
fi

echo
echo "Formatting complete."