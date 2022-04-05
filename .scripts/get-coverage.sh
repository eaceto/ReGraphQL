#! /usr/bin/env sh

GO=go
GOTEST="$GO test"
GOCOVER="$GO tool cover"


$GOTEST -v -coverprofile=docs/reports/coverage/coverage.out ./...
$GOCOVER -func=docs/reports/coverage/coverage.out -o docs/reports/coverage/coverage.funcs
$GOCOVER -html=docs/reports/coverage/coverage.out -o docs/reports/coverage/coverage.html
