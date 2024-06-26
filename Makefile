APP := canodb
BINDIR := bin
CONFDIR := config

.PHONY: all clean build linux freebsd

all: linux freebsd

clean:
	@if [ -d ${BINDIR} ]; then rm -rf ${BINDIR}/*; else exit 0; fi

build:
	go build -o ${BINDIR}/${APP} -ldflags "-s -w" .

linux:
	@# linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINDIR}/${APP} -ldflags "-s -w" .
	@cp -r ${CONFDIR}/ ${BINDIR}/ && cd ${BINDIR}/ && tar -zcf ${APP}-linux_amd64.tar.gz ${APP} ${CONFDIR}/ && rm -rf ${APP} ${CONFDIR}/ && cd ../
	@# linux-arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ${BINDIR}/${APP} -ldflags "-s -w" .
	@cp -r ${CONFDIR}/ ${BINDIR}/ && cd ${BINDIR}/ && tar -zcf ${APP}-linux_arm64.tar.gz ${APP} ${CONFDIR}/ && rm -rf ${APP} ${CONFDIR}/ && cd ../

freebsd:
	@# freebsd-amd64:
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o ${BINDIR}/${APP} -ldflags "-s -w" .
	@cp -r ${CONFDIR}/ ${BINDIR}/ && cd ${BINDIR}/ && tar -jcf ${APP}-freebsd_amd64.tar.bz2 ${APP} ${CONFDIR}/ && rm -rf ${APP} ${CONFDIR}/ && cd ../
	@# freebsd-arm64:
	CGO_ENABLED=0 GOOS=freebsd GOARCH=arm64 go build -o ${BINDIR}/${APP} -ldflags "-s -w" .
	@cp -r ${CONFDIR}/ ${BINDIR}/ && cd ${BINDIR}/ && tar -jcf ${APP}-freebsd_arm64.tar.bz2 ${APP} ${CONFDIR}/ && rm -rf ${APP} ${CONFDIR}/ && cd ../
