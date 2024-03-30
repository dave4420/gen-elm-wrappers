#!/bin/bash
set -euo pipefail

: ${BINARY_VERSION?}

rm -rf npm
mkdir npm

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

# general approach taken from <https://sentry.engineering/blog/publishing-binaries-on-npm>

list-binaries-to-build() {
    perl <targets.txt -ne 'print if not /^#/'
}

cp -r npm-main-template npm/$package_name
cp README.md npm/$package_name/README.md
printf '{' >> npm/$package_name/arch-packages.json
comma=''

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
        s390x)
            process_arch=s390
            ;;
        loong64)
            process_arch=loong64
            ;;
        s390x)
            process_arch=s390x
            ;;
        *)
            continue
    esac

    arch_package_name=@dave4420/$package_name-$process_platform-$process_arch
    mkdir -p npm/$arch_package_name/bin
    cat <<PACKAGE_JSON > npm/$arch_package_name/package.json
{
    "name": "$arch_package_name",
    "version": "$BINARY_VERSION",
    "os": [ "$process_platform" ],
    "cpu": [ "$process_arch" ]
}
PACKAGE_JSON
    cat <<README > npm/$arch_package_name/README.md
You should not need to install this package directly.

Instead,

- either install [$package_name](https://www.npmjs.com/package/$package_name)
- or manually [download the binary for your architecture](https://github.com/dave4420/gen-elm-wrappers/releases/tag/$BINARY_VERSION)
README
    cp \
        out/gen-elm-wrappers-$GOOS-$GOARCH-$BINARY_VERSION$BINARY_EXT \
        npm/$arch_package_name/bin/gen-elm-wrappers$BINARY_EXT
    printf '%s\n  "%s": "%s"' \
        "$comma" \
        "${process_platform}-${process_arch}" \
        $arch_package_name \
        >> npm/$package_name/arch-packages.json
    comma=','
done

printf '\n}\n' >> npm/$package_name/arch-packages.json

node -e '
    const pj = require("./package.json");
    pj.version = process.env.BINARY_VERSION;
    delete pj.private;
    pj.scripts = {
        "postinstall": "node ./install.js"
    };
    pj.bin = "bin/cli.js";
    delete pj.devDependencies;
    pj.optionalDependencies = {};
    fs.readdirSync("npm").forEach((dir) => {
        if (dir === "'$package_name'") {
            return;
        }
        pj.optionalDependencies[dir] = process.env.BINARY_VERSION;
    });
    process.stdout.write(JSON.stringify(pj, null, 2));
    process.stdout.write("\n");
' > npm/$package_name/package.json
