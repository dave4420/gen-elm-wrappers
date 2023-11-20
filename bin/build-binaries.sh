#!/bin/bash
set -euo pipefail

if [[ -z "$BINARY_VERSION" ]] ; then
    printf 'BINARY_VERSION not set\n' >&2
    exit 1
fi

list-binaries-to-build() {
    # DAVE: extract this from suitable json file
    echo darwin amd64
    echo darwin arm64
    echo linux amd64
    echo windows amd64
}

rm -rf out
mkdir out

list-binaries-to-build | while read GOOS GOARCH ; do
    case $GOOS in
        windows)
            BINARY_EXT=.exe
            ;;
        *)
            BINARY_EXT=''
            ;;
    esac
    BINARY_NAME="out/gen-elm-wrappers-$GOOS-$GOARCH-$BINARY_VERSION$BINARY_EXT"
    export GOOS GOARCH BINARY_EXT BINARY_NAME
    bin/build-binary.sh
done
