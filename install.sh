#!/usr/bin/env bash
set -euo pipefail

BINARY_NAME="dotback"
SOURCE_BINARY="./${BINARY_NAME}"
SYSTEM_BIN="/usr/local/bin"
USER_BIN="${HOME}/.local/bin"

if [[ ! -f "${SOURCE_BINARY}" ]]; then
  echo "error: '${SOURCE_BINARY}' not found. Build or place the binary in the project root first." >&2
  exit 1
fi

if [[ -w "${SYSTEM_BIN}" ]]; then
  DEST_DIR="${SYSTEM_BIN}"
else
  DEST_DIR="${USER_BIN}"
fi

mkdir -p "${DEST_DIR}"
install -m 0755 "${SOURCE_BINARY}" "${DEST_DIR}/${BINARY_NAME}"

echo "installed ${BINARY_NAME} to ${DEST_DIR}/${BINARY_NAME}"

