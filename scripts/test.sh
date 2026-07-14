#!/usr/bin/env bash

set -e

echo "[INFO] Running unit tests..."
go test ./tests/unit/... -v

echo "[INFO] Running integration tests..."
go test ./tests/integration/... -v

echo "[INFO] All tests completed successfully."