#!/bin/bash
set -euo pipefail

list-binaries-to-build() {
    echo darwin amd64
    echo darwin arm64
    echo linux amd64
    echo windows amd64 .exe
}

rm -rf out
mkdir out

list-binaries-to-build | while read GOOS GOARCH BINARY_EXT ; do
    BINARY_NAME="out/gen-elm-wrappers-$GOOS-$GOARCH-$BINARY_VERSION$BINARY_EXT"
    export GOOS GOARCH BINARY_EXT BINARY_NAME
    bin/build-binary.sh
done
