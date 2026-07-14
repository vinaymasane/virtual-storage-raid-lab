#!/usr/bin/env bash

set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "${ROOT}"

echo "[INFO] Collecting runtime artifacts..."

make collect                                                 
echo "[INFO] Done."