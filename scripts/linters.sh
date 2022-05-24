#!/usr/bin/env bash
echo "golang-ci lint..."
golangci-lint run ./...

echo "gogroup..."
gogroup -order std,other,prefix=GameOfLife  $(find . -type f -name "*.go" | grep -v "vendor/")