.language:.beagle/build.sh
#!/bin/bash

set -ex

git config --global --add safe.directory $PWD

mkdir -p .tmp

export CGO_ENABLED=0

# 获取当前 Git 提交和描述
export GIT_COMMIT=$(git rev-parse --short HEAD)
export GIT_DESCRIBE=$(git describe --tags --always)

# BUILD版本
BUILD_VERSION="${BUILD_VERSION:-v2.8.1}"

# 设置链接标志
LDFLAGS="-s -w -X github.com/tomwright/dasel/v2/internal.Version=${BUILD_VERSION} -X github.com/tomwright/dasel/v2/internal.Commit=${GIT_COMMIT} -X github.com/tomwright/dasel/v2/internal.GitDescribe=${GIT_DESCRIBE}"

export GOARCH=amd64
go build --ldflags "${LDFLAGS}" -o .tmp/dasel-linux-$GOARCH ./cmd/dasel

export GOARCH=arm64
go build --ldflags "${LDFLAGS}" -o .tmp/dasel-linux-$GOARCH ./cmd/dasel
