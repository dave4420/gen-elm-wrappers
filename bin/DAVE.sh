#!/bin/bash
set -euo pipefail

print_arg() {
    printf '[%d]\t%s\n' "$i" "$1" >&2
}

i=0
print_arg "$0"

for arg in "$@"; do
    i=$((i+1))
    print_arg "$arg"
done
