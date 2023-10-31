#!/bin/bash
set -euo pipefail

test_against_elm_json() {
    (
        printf 'Testing with %s\n' "$1"

        rm -rf test
        mkdir test
        cd test

        cat >elm.json

        mkdir src
        cat >src/Helpers.elm <<EOF
module Helpers exposing (..)
import Time
maybePosixFromMillis : Int -> Maybe Time.Posix
maybePosixFromMillis = Time.millisToPosix >> Just
EOF
        cat >src/Main.elm <<EOF
module Main exposing (main)
import Type.DictTimePosix
main : Program () () Never
main = Debug.todo "main"
EOF

        ../gen-elm-wrappers

        elm make src/Main.elm --debug
    )
}

go test github.com/dave4420/gen-elm-wrappers/src
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
            "elm/json": "1.1.3",
            "elm/time": "1.0.0"
        },
        "indirect": {
        }
    },
    "test-dependencies": {
        "direct": {
        },
        "indirect": {
        }
    },
    "gen-elm-wrappers": {
        "generate": [
            {
                "underlying-type": "Dict",
                "wrapper-type": "Type.DictTimePosix.DictTimePosix",
                "public-key-type": "Time.Posix",
                "private-key-type": "Int",
                "private-key-to-public-key": "Helpers.maybePosixFromMillis",
                "public-key-to-private-key": "Time.posixToMillis"
            }
        ]
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
            "elm/time": "1.0.0",
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
    },
    "gen-elm-wrappers": {
        "generate": [
            {
                "underlying-type": "Dict",
                "wrapper-type": "Type.DictTimePosix.DictTimePosix",
                "public-key-type": "Time.Posix",
                "private-key-type": "Int",
                "private-key-to-public-key": "Helpers.maybePosixFromMillis",
                "public-key-to-private-key": "Time.posixToMillis"
            }
        ]
    }
}
EOF
