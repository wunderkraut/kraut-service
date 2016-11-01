#!/bin/bash

#!/usr/bin/env bash
set -e

# We should be determining these automatically somehow?
GOVERSION="latest"

export KRAUT_PKG='github.com/wunder/kraut-service'
export KRAUT_BUILD_PATH="${KRAUT_BUILD_PATH:-/app/bin}"
export KRAUT_BINARY_NAME="kraut-service"

export KRAUT_BUILD_BINARY_PATH="${KRAUT_BUILD_PATH}/${KRAUT_BINARY_NAME}"

# Build a bundle
bundle() {
    local bundle="$1"; shift
    echo "---> Making bundle: $(basename "$bundle") (in $DEST)"
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
