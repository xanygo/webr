#!/bin/bash

go get github.com/xanygo/anygo
go mod tidy
go generate ./...