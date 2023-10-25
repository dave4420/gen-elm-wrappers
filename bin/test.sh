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
    "elm-version": "0.19.1"
}
EOF

mkdir src

cat >src/Main.elm <<EOF
module Main exposing (main)
import Data.ByDate
main = Debug.todo "main"
EOF

../gen-elm-wrappers

elm make src/Main.elm --debug
