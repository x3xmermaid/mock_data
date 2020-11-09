#!/bin/sh
echo ">>> Compiling Inventory Service into BusyBox's binary file..."
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./_bin/mock-app-data -v ./cmd/main.go
if [ $? -ne 0 ]; then
    echo "->>> Compiling Failed."
    exit -1
fi
echo "->>> Compiling Success."
echo "->>> Buidling into docker image..."
version=`cat VERSION`
if [ "$1" = "dev" ]; then
    version=$version"_dev"
fi
docker build -t x3xmermaid/mock-app-data:latest -t x3xmermaid/mock-app-data:$version -f Dockerfile-Netmonk-Mock-App-Data .
if [ $? -ne 0 ]; then
    echo "->>> Building Failed."
    exit -1
fi
echo "->>> Building Success."