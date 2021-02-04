.PHONY: gitlab/build
gitlab/build: go/mod/tidy go/version go/get go/fmt go/generate go/build ## Builds the app on Gitlab

box/blob.go:
	@if [ ! -f box/blob.go ]; \
	then $(MAKE) go/generate; \
    fi

.PHONY: gitlab/release/all
gitlab/release/all: gitlab/release/darwin/amd64 gitlab/release/solaris/amd64 gitlab/release/freebsd/386 gitlab/release/freebsd/amd64 gitlab/release/freebsd/arm gitlab/release/openbsd/386 gitlab/release/openbsd/amd64 gitlab/release/openbsd/arm gitlab/release/linux/386 gitlab/release/linux/amd64 gitlab/release/linux/arm gitlab/release/windows/386 gitlab/release/windows/amd64 ## Compile to multiple architectures

.PHONY: gitlab/release/darwin/amd64
gitlab/release/darwin/amd64: box/blob.go ## Build binary for darwin amd64
	@mkdir -p dist
	GOOS=darwin GOARCH=amd64 go build -o dist/tecli-darwin-amd64 main.go

.PHONY: gitlab/release/solaris/amd64
gitlab/release/solaris/amd64: box/blob.go ## Build binary for solaris amd64
	@mkdir -p dist
	GOOS=solaris GOARCH=amd64 go build -o dist/tecli-solaris-amd64 main.go

.PHONY: gitlab/release/freebsd/386
gitlab/release/freebsd/386: box/blob.go ## Build binary for freebsd 386
	@mkdir -p dist
	GOOS=freebsd GOARCH=386 go build -o dist/tecli-freebsd-386 main.go

.PHONY: gitlab/release/freebsd/amd64
gitlab/release/freebsd/amd64: box/blob.go ## Build binary for freebsd amd64
	@mkdir -p dist
	GOOS=freebsd GOARCH=amd64 go build -o dist/tecli-freebsd-amd64 main.go

.PHONY: gitlab/release/freebsd/arm
gitlab/release/freebsd/arm: box/blob.go ## Build binary for freebsd arm
	@mkdir -p dist
	GOOS=freebsd GOARCH=arm go build -o dist/tecli-freebsd-arm main.go

.PHONY: gitlab/release/openbsd/386
gitlab/release/openbsd/386: box/blob.go ## Build binary for openbsd 386
	@mkdir -p dist
	GOOS=openbsd GOARCH=386 go build -o dist/tecli-openbsd-386 main.go

.PHONY: gitlab/release/openbsd/amd64
gitlab/release/openbsd/amd64: box/blob.go ## Build binary for openbsd amd64
	@mkdir -p dist
	GOOS=openbsd GOARCH=amd64 go build -o dist/tecli-openbsd-amd64 main.go

.PHONY: gitlab/release/openbsd/arm
gitlab/release/openbsd/arm: box/blob.go ## Build binary for openbsd arm
	@mkdir -p dist
	GOOS=openbsd GOARCH=arm go build -o dist/tecli-openbsd-arm main.go

.PHONY: gitlab/release/linux/386
gitlab/release/linux/386: box/blob.go ## Build binary for linux 386
	@mkdir -p dist
	GOOS=linux GOARCH=386 go build -o dist/tecli-linux-386 main.go

.PHONY: gitlab/release/linux/amd64
gitlab/release/linux/amd64: box/blob.go ## Build binary for linux amd64
	@mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -o dist/tecli-linux-amd64 main.go

.PHONY: gitlab/release/linux/arm
gitlab/release/linux/arm: box/blob.go ## Build binary for linux arm
	@mkdir -p dist
	GOOS=linux GOARCH=arm go build -o dist/tecli-linux-arm main.go

.PHONY: gitlab/release/windows/386
gitlab/release/windows/386: box/blob.go ## Build binary for windows 386
	@mkdir -p dist
	GOOS=windows GOARCH=386 go build -o dist/tecli-windows-386 main.go

.PHONY: gitlab/release/windows/amd64
gitlab/release/windows/amd64: box/blob.go ## Build binary for windows amd64
	@mkdir -p dist
	GOOS=windows GOARCH=amd64 go build -o dist/tecli-windows-amd64 main.go
