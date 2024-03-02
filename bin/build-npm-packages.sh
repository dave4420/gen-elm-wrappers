#!/bin/bash
set -euo pipefail

rm -rf npm
mkdir npm

# DAVE: install gomplate

# general approach taken from <https://sentry.engineering/blog/publishing-binaries-on-npm>

# DAVE: build each arch-specific package
# DAVE: build general package
