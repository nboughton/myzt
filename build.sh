#!/bin/bash

rm -rf build
GOOS=darwin GOARCH=amd64 go build -o build/myzt.osx
GOOS=linux GOARCH=amd64 go build -o build/myzt.linux
tar czvf build/myzt.tar.gz build/myzt.*