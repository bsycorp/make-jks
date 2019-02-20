all: compile

clean:
	rm -rf build

compile: clean
	go get -d -t ./cmd/
	go test -v ./cmd/
	GOARCH=amd64 GOOS=linux go build -o ./build/make-jks-linux-x64 ./cmd/
	GOARCH=amd64 GOOS=darwin go build -o ./build/make-jks-darwin-x64 ./cmd/
	GOARCH=amd64 GOOS=windows go build -o ./build/make-jks-win-x64.exe ./cmd/
