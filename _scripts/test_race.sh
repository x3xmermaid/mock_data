#!/bin/sh
go test -v -race ./... | grep -E "PASS|FAIL|coverage"