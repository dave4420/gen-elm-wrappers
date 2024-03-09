#!/bin/bash
set -euo pipefail

: ${BINARY_VERSION?}

rm -rf npm
mkdir npm

gomplate_version=3.11.7
package_name=gen-elm-wrappers

platform="$(node -e 'console.log(process.platform)')"
case "$platform" in
    linux)
        ;;
    darwin)
        ;;
    *)
        printf 'Unsupported node process.platform: %s\n' "$platform" >&2
        exit 2
esac

arch="$(node -e 'console.log(process.arch)')"
case "$arch" in
    x64)
        arch=amd64
        ;;
    arm64)
        ;;
    *)
        printf 'Unsupported node process.arch: %s\n' "$arch" >&2
        exit 2
esac

gomplate="gomplate_${platform}-${arch}"
if ! [[ -x bin/$gomplate ]] ; then
    curl https://github.com/hairyhenderson/gomplate/releases/download/v$gomplate_version/$gomplate \
        --location \
        --continue-at - \
        --output bin/$gomplate
    chmod +x bin/$gomplate
fi

# general approach taken from <https://sentry.engineering/blog/publishing-binaries-on-npm>

list-binaries-to-build() {
    perl <targets.txt -ne 'print if not /^#/'
}

list-binaries-to-build | while read GOOS GOARCH ; do
    BINARY_EXT=''
    # keep in sync with build-binaries.sh
    case $GOOS in
        aix)
            process_platform=aix
            ;;
        darwin)
            process_platform=darwin
            ;;
        freebsd)
            process_platform=freebsd
            ;;
        linux)
            process_platform=linux
            ;;
        openbsd)
            process_platform=openbsd
            ;;
        windows)
            process_platform=win32
            BINARY_EXT=.exe
            ;;
        *)
            continue
    esac

    case $GOARCH in
        386)
            process_arch=ia32
            ;;
        amd64)
            process_arch=x64
            ;;
        arm)
            process_arch=arm
            ;;
        arm64)
            process_arch=arm64
            ;;
        ppc64|ppc64le)
            process_arch=ppc64
            ;;
        s390x)
            process_arch=s390
            ;;
        wasm)
            #DAVE
            ;;
        loong64)
            process_arch=loong64
            ;;
        mips|mipsle|mips64|mips64le)
            #DAVE
            ;;
        s390x)
            process_arch=s390x
            ;;
    esac

    arch_package_name=$package_name-$process_platform-$process_arch
    mkdir -p npm/$arch_package_name/bin
    cat <<PACKAGE_JSON > npm/$arch_package_name/package.json
{
    "name": "$arch_package_name",
    "version": "$BINARY_VERSION",
    "os": [ "$process_platform" ],
    "cpu": [ "$process_arch" ]
}
PACKAGE_JSON
    cp \
        out/gen-elm-wrappers-$GOOS-$GOARCH-$BINARY_VERSION$BINARY_EXT \
        npm/$arch_package_name/bin/gen-elm-wrappers$BINARY_EXT
done

# DAVE: build general package
