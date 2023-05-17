APP=$(shell basename $(shell git remote get-url origin))
REGISTRY=maaak7
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=linux    # windows / macos / linux
TARGETARCH=arm64  # amd64
FILE=kbot		  # kbot.exe

format:
	gofmt -s -w ./

lint:
	golint

get:
	go get

test:
	go test -v

linux:
	${MAKE} build TARGETOS=linux TARGETARCH=${TARGETARCH}

windows:
	${MAKE} build TARGETOS=windows TARGETARCH=${TARGETARCH} FILE=kbot.exe

macos:
	${MAKE} build TARGETOS=darwin TARGETARCH=${TARGETARCH}

build: format get
	CGO_ENABLED=auto GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o ${FILE} -ldflags "-X="github.com/maaak7/kbot/cmd.appVersion=${VERSION}

image:
	docker build . -t ${REGISTRY}/${APP}:${VERSION}-${TARGETARCH} \
    --build-arg CGO_ENABLED=auto \
    --build-arg TARGETARCH=${TARGETARCH} \
    --build-arg TARGETOS=${TARGETOS}

push:
	docker push ${REGISTRY}/${APP}:${VERSION}-${TARGETARCH}

clean:
	rm -f kbot || rm -f kbot.exe
	docker rmi -f ${REGISTRY}/${APP}:${VERSION}-${TARGETARCH}