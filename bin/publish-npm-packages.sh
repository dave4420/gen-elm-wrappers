#!/bin/bash
set -exuo pipefail
IFS=$'\n\t'

cd npm

for package in $(find . -type d -exec sh -c '[ -f {}/package.json ]' \; -print) ; do
    # ordered so that the package with the shortest name is published last,
    # after its dependencies
    if ! npm publish $package ; then
        printf 'Error code %d; hopefully already published?\n' $?
    fi
done
