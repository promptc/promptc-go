#!/bin/bash
go build -v -o "./out/ptc" -ldflags "-s -w" ./cli/*.go