#!/bin/bash
set -euo pipefail

if [[ -z "$BINARY_NAME" ]] ; then
    printf 'BINARY_NAME not set\n' >&2
    exit 1
fi

if [[ -z "$BINARY_VERSION" ]] ; then
    printf 'BINARY_VERSION not set\n' >&2
    exit 1
fi

go build \
    -o "$BINARY_NAME" \
    -ldflags "-X main.Version=$BINARY_VERSION" \
    github.com/dave4420/gen-elm-wrappers/src
