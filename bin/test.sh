#!/bin/bash
set -euo pipefail

rm -rf test

go build .

mkdir test
cd test

cat >elm.json <<EOF
{
    "type": "application",
    "source-directories": [
        "src"
    ],
    "elm-version": "0.19.1",
    "dependencies": {
        "direct": {
            "elm/core": "1.0.5",
            "elm/json": "1.1.3"
        },
        "indirect": {
        }
    },
    "test-dependencies": {
        "direct": {
        },
        "indirect": {
        }
    }
}
EOF

mkdir src

cat >src/Main.elm <<EOF
module Main exposing (main)
import Type.DictInt
main = Debug.todo "main"
EOF

../gen-elm-wrappers

elm make src/Main.elm --debug
