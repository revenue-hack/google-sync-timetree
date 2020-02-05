.PHONY: build clean deploy

FUNC:=treesync

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o 'bin/${FUNC}' sync/*

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

invoke: build
	sls invoke local -f ${FUNC}
