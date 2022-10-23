#!/bin/sh
APP="opio"

# OS별 필드
GOOS=linux GOARCH=amd64 go build -o ./bin/linux/${APP} opio.go stereo.go linux.go macos.go pathfunc.go
GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin/${APP} opio.go stereo.go linux.go macos.go pathfunc.go

# Github Release에 업로드 하기위해 압축
cd ./bin/linux/ && cp ../../../../install/linux/install_CentOS* . && tar -zcvf ../${APP}_linux_x86-64.tgz . && cd -
cd ./bin/darwin/ && cp -r ../../../../install/macos/opio.app . && tar -zcvf ../${APP}_darwin_x86-64.tgz . && cd -

# 삭제
rm -rf ./bin/linux
rm -rf ./bin/darwin
