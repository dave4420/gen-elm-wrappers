#!/bin/bash
set -euo pipefail

: ${BINARY_VERSION?}

list-binaries-to-build() {
    perl <targets.txt -ne 'print if not /^#/'
}

rm -rf out
mkdir out

list-binaries-to-build | while read GOOS GOARCH ; do
    case $GOOS in
        android|ios)
            # get an error when building for android
            # get a warning when building for ios, suggests it won't work
            # I'm not expecting many people to want to run this software
            # on these platforms
            continue
            ;;
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
