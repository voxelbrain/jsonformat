#!/bin/bash

PLATFORMS="darwin/amd64 linux/amd64"
APPNAME=$(basename `pwd`)
VERSION=$(cat README.md | tail -n4 | grep -i Version | cut -d' ' -f2)

for platform in $PLATFORMS; do
	export GOOS=$(echo $platform | cut -d/ -f1)
	export GOARCH=$(echo $platform | cut -d/ -f2)
	BINNAME="${APPNAME}_${VERSION}_${GOOS}_${GOARCH}"
	echo "Building $BINNAME..."
	CGO_ENABLED=0 go build -o $BINNAME *.go
done
