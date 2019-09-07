#!/usr/bin/env bash
#
# This script tests the application from source.

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

cd "$DIR"

rm -rf build/e2e

BINARY="${BINARY:-go run cmd/seedr/main.go}"

$BINARY generate --seed test/testdata/seed-e2e/full --target build/e2e/full
$BINARY generate --seed test/testdata/seed-e2e/noseedfile --target build/e2e/noseedfile
