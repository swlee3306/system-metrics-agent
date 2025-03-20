#!/bin/bash

# 사용 방법: GOARCH=amd64 ./build.sh 

# 환경 변수 설정
GOOS=linux
GOARCH=${GOARCH:-arm64}  # 기본값을 arm64로 설정 (환경 변수가 없을 경우)
CGO_ENABLED=0

# 빌드 정보 설정
BUILD_DATE=$(date "+%Y-%m-%d %H:%M:%S")
BUILD_REVISION=$(git rev-parse HEAD)
BUILD_REVISION_SHORT=$(git rev-parse --short HEAD)
BUILD_BRANCH=$(git rev-parse --abbrev-ref HEAD)
BUILD_TAG=$(git describe --tags --abbrev=0 2>/dev/null || echo "none")

# 빌드 정보 출력
echo "🔨 Building for GOOS=$GOOS, GOARCH=$GOARCH"

# 환경 변수 설정 (Mac에서 Linux aarch64용 빌드)
GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=$CGO_ENABLED \
go build -ldflags "\
    -X 'main._buildDate=${BUILD_DATE}' \
    -X 'main._buildRevision=${BUILD_REVISION}' \
    -X 'main._buildRevisionShort=${BUILD_REVISION_SHORT}' \
    -X 'main._buildBranch=${BUILD_BRANCH}' \
    -X 'main._buildTag=${BUILD_TAG}'" \
    -o system-agent-linux-$GOARCH ../main.go

echo "✅ 빌드 완료: system-agent-linux-arm64"