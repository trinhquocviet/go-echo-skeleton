#!/bin/bash
# Read more: https://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5
echo "Start build linux"
echo "==="
echo "Remove all old file"
rm -rf ../build
mkdir ../build
echo "==="
echo "Building linux amd64"
GOOS=linux GOARCH=amd64 go build -o ../build/main ../main.go

echo "Copy files"
cp -R ../config ../build/config
cp ../config/.env ./build/config/.env
echo "Allmost done!!!"