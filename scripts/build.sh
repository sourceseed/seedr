#!/usr/bin/env bash
#
# This script builds the application from source for multiple platforms.

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

cd "$DIR"

# Determine the arch/os combos we're building for
XC_ARCH=${XC_ARCH:-"386 amd64 arm"}
XC_OS=${XC_OS:-linux darwin windows}
XC_EXCLUDE_OSARCH="!darwin/arm !darwin/386"


# Delete the build dir
echo "==> Removing build directory..."
rm -rf build/*
mkdir -p build/{bin,pkg}

if ! which gox > /dev/null; then
    echo "==> Installing gox..."
    go get -u github.com/mitchellh/gox
fi

# Instruct gox to build statically linked binaries
export CGO_ENABLED=0

LD_FLAGS="-s -w "
LD_FLAGS+="-X 'github.com/NoUseFreak/go-vembed.version=`git describe --tags --abbrev=0`' "
LD_FLAGS+="-X 'github.com/NoUseFreak/go-vembed.buildDate=`date +%Y-%m-%d\ %H:%M`' "
LD_FLAGS+="-X 'github.com/NoUseFreak/go-vembed.gitCommit=`git rev-parse --short HEAD`' "
LD_FLAGS+="-X 'github.com/NoUseFreak/go-vembed.gitState=`git diff --quiet || echo 'dirty'`' "
LD_FLAGS+="-X 'github.com/NoUseFreak/go-vembed.gitBranch=`git symbolic-ref -q --short HEAD`' "
LD_FLAGS+="-X 'github.com/NoUseFreak/go-vembed.gitSummary=`git describe --tags --dirty --always`' "

# Ensure all remote modules are downloaded and cached before build so that
# the concurrent builds launched by gox won't race to redundantly download them.
go mod download

# Build!
echo "==> Building..."
gox \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -osarch="${XC_EXCLUDE_OSARCH}" \
    -ldflags "${LD_FLAGS}" \
    -output "build/bin/{{.OS}}_{{.Arch}}/${PWD##*/}" \
    .

# Zip and copy to the dist dir
echo "==> Packaging..."
for PLATFORM in $(find ./build/bin -mindepth 1 -maxdepth 1 -type d); do
    OSARCH=$(basename ${PLATFORM})
    echo "--> ${OSARCH}"

    pushd $PLATFORM >/dev/null 2>&1
    zip ../../pkg/${OSARCH}.zip ./*
    popd >/dev/null 2>&1
done

# Done!
echo
echo "==> Results:"
du -ah build/*