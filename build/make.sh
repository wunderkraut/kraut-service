#!/bin/bash

#!/usr/bin/env bash
set -e

# We should be determining these automatically somehow?
GOVERSION="latest"

export KRAUT_PKG='github.com/wunderkraut/radi-service'
export KRAUT_BUILD_PATH="${KRAUT_BUILD_PATH:-./bin}"
export KRAUT_BINARY_NAME="${KRAUT_BINARY_NAME:-radi-service}"

export KRAUT_BUILD_BINARY_PATH="${KRAUT_BUILD_PATH}/${KRAUT_BINARY_NAME}"

# Build a bundle
bundle() {
    local bundle="$1"; shift
    echo "---> Make-bundle: $(basename "$bundle") (in $DEST)"
    source "build/$bundle" "$@"
}

if [ $# -gt 0 ]; then
    bundles=($@)
    for bundle in ${bundles[@]}; do
        export DEST=.
        ABS_DEST="$(cd "$DEST" && pwd -P)"
        bundle "$bundle"
        echo
    done
fi
