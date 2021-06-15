build:
	@echo "Building binary file for Linux..."
	GOARCH=amd64 GOOS=linux go build -o bin/spb-linux-amd64
	@echo "Building binary file for Windows..."
	GOARCH=amd64 GOOS=windows go build -o bin/spb-win-amd64.exe
	@echo "Building binary file for MacOS..."
	GOARCH=amd64 GOOS=darwin go build -o bin/spb-darwin-amd64
	GOARCH=arm64 GOOS=darwin go build -o bin/spb-darwin-arm64