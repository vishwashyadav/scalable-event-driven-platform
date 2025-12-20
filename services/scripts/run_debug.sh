#!/usr/bin/env bash
# Helper to run the auto-reload debug loop locally.
# Requires: CompileDaemon (github.com/githubnemo/CompileDaemon) and Delve (dlv).

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT_DIR"

TMP_BIN="./tmp/debug_bin"
mkdir -p ./tmp

echo "Starting CompileDaemon -> Delve (headless)"
echo "Make sure you installed:"
echo "  go install github.com/githubnemo/CompileDaemon@latest"
echo "  brew install go-delve/delve/delve   # or go install github.com/go-delve/delve/cmd/dlv@latest"

CompileDaemon -build="go build -o ${TMP_BIN} ./api-service/cmd/server" -command="dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ${TMP_BIN}"
