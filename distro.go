package main

import (
	"errors"
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
	Void    OS_ID = "void"
	Fedora  OS_ID = "fedora"
	Windows OS_ID = "windows"
	Linux   OS_ID = "linux"
	Darwin  OS_ID = "darwin"
	Unknown OS_ID = "unknown"
)

func OSRich() string {
	switch runtime.GOOS {
	// Linux
	case "linux":
		cfg, err := ini.Load("/etc/os-release")
		if err != nil {
			return "Unknown Linux"
		}
		return cfg.Section("").Key("PRETTY_NAME").String()

	// BSD
	case "openbsd":
		return "OpenBSD"

	case "freebsd":
		return "FreeBSD"

	case "netbsd":
		return "NetBSD"

	case "dragonflybsd":
		return "DragonFly BSD"

	// Unix
	case "solaris":
		return "Solaris"

	case "darwin":
		return "Apple MacOS"

	// Windows and other
	case "windows":
		return "Microsoft Windows"

	default:
		return "Unknown OS (" + runtime.GOOS + ")"
	}
}

func OSType() OS_ID {
	if runtime.GOOS == "linux" {
		cfg, err := ini.Load("/etc/os-release")
		if err != nil {
			return Linux
		}
		id := cfg.Section("").Key("ID").String()
		switch id {
		// Fedora
		case "fedora", "centos":
			return Fedora

		// Debian
		case "debian", "ubuntu", "linuxmint":
			return Debian

		// Arch Linux
		case "arch", "artix", "manjaro":
			return Arch

		// Void Linux
		case "void":
			return Void

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

func sysupdate(system OS_ID, root bool, cross bool) error {
	if cross {
		kageroExists := utils.CommandExists("kagero")
		kazeExists := utils.CommandExists("kaze")
		snapExists := utils.CommandExists("snap")
		if kageroExists || kazeExists || snapExists {
			utils.Printer.Println("Updating cross-platform packages...", pringo.CyanBright)
		}

		// Cross-platform package managers
		if kageroExists {
			systemupdate.Kagero(false)
		}
		if kazeExists {
			systemupdate.Kaze(root)
		}
		if snapExists {
			systemupdate.Snap(root)
		}

		// Ganyu tool via gosdk
		location, err := os.Executable()
		if err != nil {
			location = os.Args[0]
		}
		if utils.CommandExists("go") && utils.IsInGopath(location) {
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
		return systemupdate.Apt(root)
	case Arch:
		return systemupdate.Arch(root)
	case Void:
		return systemupdate.Void(root)
	case Fedora:
		return systemupdate.Fedora(root)
	case Darwin:
		if utils.CommandExists("brew") {
			return systemupdate.Brew(root)
		} else {
			utils.Printer.Errorln("No package manager found!", pringo.Red)
			return nil
		}
	case Windows:
		if utils.CommandExists("choco") {
			return systemupdate.Choco(root)
		} else {
			utils.Printer.Errorln("No package manager found!", pringo.Red)
			return nil
		}
	default:
		return ErrNotSupported
	}
}

var ErrNotSupported = errors.New("system not supported")
