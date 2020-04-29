#!/bin/sh

PATH=$PATH:$GOPATH/bin protoc -I=.. --go_out=. --go-grpc_out=. ../*.proto
rm -rf gauge_messages
mv github.com/getgauge/gauge-proto/go/gauge_messages ./
rm -rf github.com
cd gauge_messages
go mod init github.com/getgauge/gauge-proto/go/gauge_messages
go build ./...
