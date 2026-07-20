#!/usr/bin/env bash

set -euo pipefail

echo "==> Cleaning generated artifacts"

rm -rf bin/

rm -rf dist/

rm -rf coverage/

rm -rf .cache/

rm -rf web/dist/

echo
echo "Workspace cleaned."