package main

import (
	"runtime"

	"github.com/go-ini/ini"
)

type OS_ID string

const (
	Debian  OS_ID = "debian"
	Arch    OS_ID = "arch"
	Fedora  OS_ID = "fedora"
	Windows OS_ID = "windows"
	Linux   OS_ID = "linux"
	Darwin  OS_ID = "darwin"
	Unknown OS_ID = "unknown"
)

func OSType() OS_ID {
	if runtime.GOOS == "linux" {
		cfg, err := ini.Load("/etc/os-release")
		if err != nil {
			return Linux
		}
		id := cfg.Section("").Key("ID").String()
		switch id {
		// Fedora
		case "centos":
		case "fedora":
			return Fedora
		// Debian
		case "ubuntu":
		case "debian":
			return Debian
		// Arch Linux
		case "artix":
		case "arch":
			return Arch
		// Everything else
		default:
			return Linux
		}
	} else if runtime.GOOS == "windows" {
		return Windows
	} else if runtime.GOOS == "darwin" {
		return Darwin
	}
	return Unknown
}
