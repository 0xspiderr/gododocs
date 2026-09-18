BINARY_NAME = gododocs

build:
	mkdir -p builds
	GOARCH=amd64 GOOS=linux go build -o builds/${BINARY_NAME}-linux ./cmd/gododocs.go
	GOARCH=amd64 GOOS=windows go build -o builds/${BINARY_NAME}-win.exe ./cmd/gododocs.go
	GOARCH=amd64 GOOS=darwin go build -o builds/${BINARY_NAME}-mac ./cmd/gododocs.go

clean:
	go clean
	rm -rf builds/
