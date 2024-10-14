build:
	@go build -o bin/terminal-games

build-all:
	@mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -o bin/terminal-games-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o bin/terminal-games-darwin-amd64
	GOOS=windows GOARCH=amd64 go build -o bin/terminal-games-windows-amd64.exe
	GOOS=linux GOARCH=arm64 go build -o bin/terminal-games-linux-arm64
	GOOS=linux GOARCH=arm go build -o bin/terminal-games-linux-arm
	GOOS=darwin GOARCH=arm64 go build -o bin/terminal-games-darwin-arm64

run: build
	@./bin/terminal-games

test:
	@go test ./...