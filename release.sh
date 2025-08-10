#!/bin/bash

go build -o ./bloggo cmd/blog/main.go

ZIP_FILE="bloggo-$(date --iso-8601=seconds).tar.gz"

tar -czvf "$ZIP_FILE" ./bloggo ./web ./scripts

rm ./bloggo

echo "$ZIP_FILE is ready!"

