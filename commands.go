package main

import (
	"github.com/Stridsvagn69420/ganyu/systemupdate"
	"github.com/Stridsvagn69420/ganyu/utils"
	"github.com/Stridsvagn69420/pringo"
)

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
		if utils.CommandExists("go") {
			utils.RunShell(false, "go", "install", "github.com/Stridsvagn69420/ganyu")
		}
	}

	// System package managers
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
