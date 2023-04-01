#!/bin/bash

go build -v -o "./out/promptc" -ldflags "-s -w" ./cli/*.go
sudo cp ./out/promptc /usr/local/bin/