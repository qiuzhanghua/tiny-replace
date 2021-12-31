all: pre build build-linux build-darwin build-windows post

pre:
	autotag write

build:
	go build -o Tiny-replace/Tiny-replace

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o Tiny-replace/linux_amd64/Tiny-replace
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o Tiny-replace/linux_386/Tiny-replace

build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o Tiny-replace/darwin_amd64/Tiny-replace
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o Tiny-replace/darwin_arm64/Tiny-replace

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o Tiny-replace/windows_amd64/Tiny-replace.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o Tiny-replace/windows_386/Tiny-replace.exe

post:
	git restore autotag.go

clean:
	rm -rf Tiny-replace
