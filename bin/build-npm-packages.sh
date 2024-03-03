#!/bin/bash
set -euo pipefail

rm -rf npm
mkdir npm

gomplate_version=3.11.7

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

# DAVE: build each arch-specific package
# DAVE: build general package
