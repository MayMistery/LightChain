# Makefile for cross-compiling the redsv project

export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
BINARY=lchain
VERSION=1.0.0
BUILD=`git rev-parse HEAD`

LDFLAGS := -s -w -X main.Version=${VERSION} -X main.Build=${BUILD}

os-archs=darwin:arm64
#os-archs=darwin:amd64 darwin:arm64 freebsd:386 freebsd:amd64 linux:386 linux:amd64 linux:arm linux:arm64 windows:386 windows:amd64 windows:arm64 linux:mips64 linux:mips64le linux:mips:softfloat linux:mipsle:softfloat linux:riscv64

all: build

build: app

app:
	@$(foreach n, $(os-archs),\
		os=$(shell echo "$(n)" | cut -d : -f 1);\
		arch=$(shell echo "$(n)" | cut -d : -f 2);\
		gomips=$(shell echo "$(n)" | cut -d : -f 3);\
		target_suffix=$${os}_$${arch};\
		echo "Build $${os}-$${arch}...";\
		env CGO_ENABLED=0 GOOS=$${os} GOARCH=$${arch} GOMIPS=$${gomips} go build -trimpath -ldflags "$(LDFLAGS)" -o ./release/${BINARY}_$${target_suffix} ;\
		echo "Build $${os}-$${arch} done";\
	)
#	@mv ./release/${BINARY}_windows_386 ./release/${BINARY}_windows_386.exe
#	@mv ./release/${BINARY}_windows_amd64 ./release/${BINARY}_windows_amd64.exe
#	@mv ./release/${BINARY}_windows_arm64 ./release/${BINARY}_windows_arm64.exe

clean:
	rm -rf ./release/