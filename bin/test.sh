#!/bin/bash
set -euo pipefail

test_against_elm_json() {
    printf 'Testing with %s\n' "$1"

    rm -rf test
    mkdir test
    cd test

    cat >elm.json

    mkdir src
    cat >src/Main.elm <<EOF
module Main exposing (main)
import Type.DictInt
main : Program () () Never
main = Debug.todo "main"
EOF

    ../gen-elm-wrappers

    elm make src/Main.elm --debug
}

go build -o gen-elm-wrappers github.com/dave4420/gen-elm-wrappers/src

test_against_elm_json 'core only' <<EOF
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

test_against_elm_json 'dict-extra included' <<EOF
{
    "type": "application",
    "source-directories": [
        "src"
    ],
    "elm-version": "0.19.1",
    "dependencies": {
        "direct": {
            "elm/core": "1.0.5",
            "elm/json": "1.1.3",
            "elm-community/dict-extra": "2.4.0"
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
