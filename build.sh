#!/bin/bash
go build -v -o "./out/promptc" -ldflags "-s -w" ./cli/*.go
GOARCH=amd64 go build -v -o "./out/promptc-amd64" -ldflags "-s -w" ./cli/*.go