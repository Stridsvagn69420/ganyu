build:
	go build -o ./build/ ganyu.go

install:
	go install -ldflags="-s -w" ganyu.go

release:
	pwsh Release.ps1
