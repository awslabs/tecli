.PHONY: gitlab/build
gitlab/build: go/mod/tidy go/version go/get go/fmt go/generate go/build ## Builds the app on Gitlab

.PHONY: gitlab/release
gitlab/release: ## Compile to multiple architectures
	@mkdir -p dist
	@echo "Compiling for every OS and Platform"
	GOOS=darwin GOARCH=amd64 go build -o dist/tecli-darwin-amd64 main.go
	GOOS=solaris GOARCH=amd64 go build -o dist/tecli-solaris-amd64 main.go

	GOOS=freebsd GOARCH=386 go build -o dist/tecli-freebsd-386 main.go
	GOOS=freebsd GOARCH=amd64 go build -o dist/tecli-freebsd-amd64 main.go
	GOOS=freebsd GOARCH=arm go build -o dist/tecli-freebsd-arm main.go

	GOOS=openbsd GOARCH=386 go build -o dist/tecli-openbsd-386 main.go
	GOOS=openbsd GOARCH=amd64 go build -o dist/tecli-openbsd-amd64 main.go
	GOOS=openbsd GOARCH=arm go build -o dist/tecli-openbsd-arm main.go

	GOOS=linux GOARCH=386 go build -o dist/tecli-linux-386 main.go
	GOOS=linux GOARCH=amd64 go build -o dist/tecli-linux-amd64 main.go
	GOOS=linux GOARCH=arm go build -o dist/tecli-linux-arm main.go

	GOOS=windows GOARCH=386 go build -o dist/tecli-windows-386 main.go
	GOOS=windows GOARCH=amd64 go build -o dist/tecli-windows-amd64 main.go