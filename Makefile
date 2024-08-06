# canodb - Makefile
# Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
#
# Use of this source code is governed by a BSD 2-Clause License that can be
# found in the LICENSE file.

APP := canodb-cli
BINDIR := bin

.PHONY: all clean build linux freebsd

all: linux freebsd

clean:
	@if [ -d ${BINDIR} ]; then rm -rf ${BINDIR}/*; else exit 0; fi

build:
	go build -o ${BINDIR}/${APP} -ldflags "-s -w" .

linux:
	@# linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINDIR}/${APP} -ldflags "-s -w" .
	@cp -r ${BINDIR}/ && cd ${BINDIR}/ && tar -zcf ${APP}-linux_amd64.tar.gz ${APP} && rm -rf ${APP} && cd ../
	@# linux-arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ${BINDIR}/${APP} -ldflags "-s -w" .
	@cp -r ${BINDIR}/ && cd ${BINDIR}/ && tar -zcf ${APP}-linux_arm64.tar.gz ${APP} && rm -rf ${APP} && cd ../

freebsd:
	@# freebsd-amd64:
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o ${BINDIR}/${APP} -ldflags "-s -w" .
	@cp -r ${BINDIR}/ && cd ${BINDIR}/ && tar -jcf ${APP}-freebsd_amd64.tar.bz2 ${APP} && rm -rf ${APP} && cd ../
	@# freebsd-arm64:
	CGO_ENABLED=0 GOOS=freebsd GOARCH=arm64 go build -o ${BINDIR}/${APP} -ldflags "-s -w" .
	@cp -r ${BINDIR}/ && cd ${BINDIR}/ && tar -jcf ${APP}-freebsd_arm64.tar.bz2 ${APP} && rm -rf ${APP} && cd ../
