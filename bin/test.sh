#!/bin/bash
set -euo pipefail

expect_files_to_contain_current_year() {
    printf 'Expecting current year to be present in: %s\n' "$*"
    year="$(date +'%Y')"
    for file in "$@" ; do
        if ! grep -qE '\b'"$year"'\b' "$file" ; then
            printf 'Current year "%s" not found in %s\n' "$year" "$file" >&2
            exit 1
        fi
    done
}

expect_success() {
    (
        printf 'Expecting success with %s\n' "$1"

        rm -rf test
        mkdir test
        cd test

        printf '%s' "$2" >elm.json

        printf '%s' "$3" >gen-elm-wrappers.json

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

expect_failure_to_generate() {
    (
        printf 'Expecting failure to generate with %s\n' "$1"

        rm -rf test
        mkdir test
        cd test

        printf '%s' "$2" >elm.json

        printf '%s' "$3" >gen-elm-wrappers.json

        ! ../gen-elm-wrappers
    )
}

elm_json_core_only='
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
    }
}
'

elm_json_with_extras='
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
    }
}
'

elm_json_with_far_future_elm_core='
{
    "type": "application",
    "source-directories": [
        "src"
    ],
    "elm-version": "0.19.1",
    "dependencies": {
        "direct": {
            "elm/core": "999.0.0",
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
    }
}
'

elm_json_with_v1_dict_extra='
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
            "elm-community/dict-extra": "1.5.0"
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
'

gen_elm_wrappers_json='
{
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
'

go test github.com/dave4420/gen-elm-wrappers/src
BINARY_VERSION='?.?.?' bin/build-binary.sh

expect_success 'core only' "$elm_json_core_only" "$gen_elm_wrappers_json"

expect_success 'dict-extra included' "$elm_json_with_extras" "$gen_elm_wrappers_json"

expect_failure_to_generate 'far future elm/core' "$elm_json_with_far_future_elm_core" "$gen_elm_wrappers_json"

expect_failure_to_generate 'v1 dict-extra' "$elm_json_with_v1_dict_extra" "$gen_elm_wrappers_json"

expect_files_to_contain_current_year LICENSE

echo PASS
