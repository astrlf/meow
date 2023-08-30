BIN=meow

run:
	go run ./

build: 
	GOARCH=amd64 GOOS=darwin  go build -o ${BIN}-darwin  ./
	GOARCH=amd64 GOOS=linux   go build -o ${BIN}-linux   ./
	GOARCH=amd64 GOOS=windows go build -o ${BIN}-windows ./

release: build
	tar  -czvf 		 ${BIN}-darwin.tar.gz ${BIN}-darwin
	tar  -czvf 		 ${BIN}-linux.tar.gz 	${BIN}-linux
	7z a -qq -tzip ${BIN}-windows.zip 	${BIN}-windows

clean:
	go clean

	rm -f ${BIN}-darwin \
				${BIN}-linux  \
				${BIN}-windows

	rm -f ${BIN}-darwin.tar.gz \
				${BIN}-linux.tar.gz  \
				${BIN}-windows.zip
