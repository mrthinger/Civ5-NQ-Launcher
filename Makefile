.DEFAULT_GOAL := build_windows64

build_windows64:
	set GOOS=windows
	set GOARCH=amd64
	go build -o bin/NQLauncher.exe -ldflags="-s -w" ./cmd/client
	upx --best bin/NQLauncher.exe