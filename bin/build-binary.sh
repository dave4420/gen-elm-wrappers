#!/bin/bash
set -euo pipefail

: ${BINARY_NAME?}
: ${BINARY_VERSION?}

set -x

go build \
    -o "$BINARY_NAME" \
    -ldflags "-X main.Version=$BINARY_VERSION" \
    github.com/dave4420/gen-elm-wrappers/src
