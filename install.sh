#!/bin/bash

go build -v -o "./out/promptc" -ldflags "-s -w" ./cli/*.go
sudo rm /usr/local/bin/promptc
sudo cp ./out/promptc /usr/local/bin