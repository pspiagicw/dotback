#!/usr/bin/env bash
set -euo pipefail

BINARY_NAME="dotback"
REPO="${REPO:-pspiagicw/dotback}"
VERSION="${VERSION:-latest}"
SYSTEM_BIN="/usr/local/bin"
USER_BIN="${HOME}/.local/bin"
TMP_DIR="$(mktemp -d)"

cleanup() {
  rm -rf "${TMP_DIR}"
}
trap cleanup EXIT

require_cmd() {
  if ! command -v "$1" >/dev/null 2>&1; then
    echo "error: required command not found: $1" >&2
    exit 1
  fi
}

pick_downloader() {
  if command -v curl >/dev/null 2>&1; then
    echo "curl"
    return
  fi
  if command -v wget >/dev/null 2>&1; then
    echo "wget"
    return
  fi
  echo "error: either 'curl' or 'wget' is required" >&2
  exit 1
}

fetch() {
  local url="$1"
  local output="$2"
  if [[ "${DOWNLOADER}" == "curl" ]]; then
    curl -fsSL "${url}" -o "${output}"
  else
    wget -qO "${output}" "${url}"
  fi
}

detect_os() {
  case "$(uname -s | tr '[:upper:]' '[:lower:]')" in
    linux*) echo "linux" ;;
    darwin*) echo "darwin" ;;
    *) echo "error: unsupported OS ($(uname -s))" >&2; exit 1 ;;
  esac
}

detect_arch() {
  case "$(uname -m | tr '[:upper:]' '[:lower:]')" in
    x86_64|amd64) echo "amd64" ;;
    aarch64|arm64) echo "arm64" ;;
    *) echo "error: unsupported architecture ($(uname -m))" >&2; exit 1 ;;
  esac
}

latest_release_api() {
  if [[ "${VERSION}" == "latest" ]]; then
    echo "https://api.github.com/repos/${REPO}/releases/latest"
  else
    echo "https://api.github.com/repos/${REPO}/releases/tags/${VERSION}"
  fi
}

select_asset_url() {
  local api_url="$1"
  local json_file="${TMP_DIR}/release.json"
  fetch "${api_url}" "${json_file}"

  local urls_file="${TMP_DIR}/urls.txt"
  grep -Eo '"browser_download_url":[[:space:]]*"[^"]+"' "${json_file}" | sed -E 's/.*"([^"]+)"/\1/' > "${urls_file}"

  local candidate=""
  while IFS= read -r url; do
    local lower
    lower="$(printf '%s' "${url}" | tr '[:upper:]' '[:lower:]')"
    if [[ "${lower}" != *"${BINARY_NAME}"* ]]; then
      continue
    fi
    if [[ "${lower}" == *"${OS}-${ARCH}"* || "${lower}" == *"${OS}_${ARCH}"* || "${lower}" == *"${ARCH}-${OS}"* || "${lower}" == *"${ARCH}_${OS}"* ]]; then
      candidate="${url}"
      break
    fi
  done < "${urls_file}"

  if [[ -z "${candidate}" ]]; then
    while IFS= read -r url; do
      local lower
      lower="$(printf '%s' "${url}" | tr '[:upper:]' '[:lower:]')"
      if [[ "${lower}" == *"${BINARY_NAME}"* && "${lower}" == *"${OS}"* && "${lower}" == *"${ARCH}"* ]]; then
        candidate="${url}"
        break
      fi
    done < "${urls_file}"
  fi

  if [[ -z "${candidate}" ]]; then
    echo "error: could not find a release asset for ${OS}/${ARCH} in ${REPO} (${VERSION})" >&2
    exit 1
  fi

  printf '%s\n' "${candidate}"
}

extract_binary() {
  local asset_path="$1"
  local binary_path=""
  case "${asset_path}" in
    *.tar.gz|*.tgz)
      require_cmd tar
      tar -xzf "${asset_path}" -C "${TMP_DIR}"
      ;;
    *.zip)
      require_cmd unzip
      unzip -q "${asset_path}" -d "${TMP_DIR}"
      ;;
    *)
      chmod +x "${asset_path}" || true
      binary_path="${asset_path}"
      ;;
  esac

  if [[ -z "${binary_path}" ]]; then
    binary_path="$(find "${TMP_DIR}" -type f -name "${BINARY_NAME}" | head -n 1 || true)"
  fi

  if [[ -z "${binary_path}" ]]; then
    echo "error: could not locate '${BINARY_NAME}' inside downloaded asset" >&2
    exit 1
  fi

  printf '%s\n' "${binary_path}"
}

DOWNLOADER="$(pick_downloader)"
OS="$(detect_os)"
ARCH="$(detect_arch)"

if [[ -w "${SYSTEM_BIN}" ]]; then
  DEST_DIR="${SYSTEM_BIN}"
else
  DEST_DIR="${USER_BIN}"
fi

ASSET_URL="$(select_asset_url "$(latest_release_api)")"
ASSET_FILE="${TMP_DIR}/asset"

echo "downloading ${REPO} (${VERSION}) for ${OS}/${ARCH}"
fetch "${ASSET_URL}" "${ASSET_FILE}"
DOWNLOADED_BINARY="$(extract_binary "${ASSET_FILE}")"

mkdir -p "${DEST_DIR}"
install -m 0755 "${DOWNLOADED_BINARY}" "${DEST_DIR}/${BINARY_NAME}"

echo "installed ${BINARY_NAME} to ${DEST_DIR}/${BINARY_NAME}"
