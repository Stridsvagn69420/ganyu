files = ganyu.go config.go distro.go meta.go help.go

compile:
	go build -o ./build/ $(files)

install:
	go install -ldflags="-s -w" $(files)

release:
	pwsh Release.ps1
