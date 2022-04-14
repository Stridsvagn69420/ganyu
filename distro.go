package main

import (
	"runtime"

	"os"

	"github.com/go-ini/ini"

	"github.com/Stridsvagn69420/ganyu/systemupdate"
	"github.com/Stridsvagn69420/ganyu/utils"
	"github.com/Stridsvagn69420/pringo"
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

func sysupdate(system OS_ID, root bool, cross bool) {
	if cross {
		// Cross-platform package managers
		if utils.CommandExists("kagero") {
			systemupdate.Kagero(false)
		}
		if utils.CommandExists("kaze") {
			systemupdate.Kaze(root)
		}
		if utils.CommandExists("snap") {
			systemupdate.Snap(root)
		}
		// Ganyu tool via gosdk
		if utils.CommandExists("go") && utils.IsInGopath(os.Args[0]) {
			utils.Printer.Println("Updating Ganyu...", pringo.CyanBright)
			err := utils.RunShell(false, "go", "install", "github.com/Stridsvagn69420/ganyu@latest")
			if err != nil {
				utils.Printer.Errorln("Failed to update Ganyu with Go SDK!", pringo.Red)
			}
		}
	}

	// System package managers
	utils.Printer.Println("Updating system packages...", pringo.CyanBright)
	switch system {
	case Debian:
		systemupdate.Apt(root)
	case Arch:
		systemupdate.Arch(root)
	case Fedora:
		systemupdate.Fedora(root)
	case Windows:
		if utils.CommandExists("choco") {
			systemupdate.Choco(root)
		} else {
			utils.Printer.Errorln("No package manager found!", pringo.Red)
		}
	default:
		utils.Printer.Errorln("Your system is currently not supported!", pringo.Red)
	}
}
