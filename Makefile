BIN=meow

run:
	go run ./

build: test
	GOARCH=amd64 GOOS=darwin  go build -o ${BIN}-darwin main.go
	GOARCH=amd64 GOOS=linux   go build -o ${BIN}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${BIN}-windows main.go

release: build
	tar -czvf ${BIN}-darwin.tar.gz ${BIN}-darwin
	tar -czvf ${BIN}-linux.tar.gz ${BIN}-linux
	7z a -tzip ${BIN}-windows.zip ${BIN}-windows.exe

test:
	go test -v ./...

clean:
	go clean

	rm -f ${BIN}-darwin \
				${BIN}-linux \
				${BIN}-windows