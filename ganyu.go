package main

import (
	"fmt"
	"os"

	"github.com/Stridsvagn69420/ganyu/systemupdate"
	"github.com/Stridsvagn69420/ganyu/utils"
)

func main() {
	// Initialize command-line flags and read config
	config := config()
	flags(&config)

	// Get command
	switch os.Args[0] {
	case "sysupdate":
		sysupdate(OSType(), config.Sysupdate.Root, config.Sysupdate.CrossPkg)

	default:
		fmt.Println("Command not found or doesn't exist yet!")
	}

	os.Exit(0)
}

func sysupdate(system OS_ID, root bool, cross bool) {
	if cross {
		// Cross-platform package managers
		if utils.CommandExists("kagero") {
			systemupdate.Kagero(root)
		}
		if utils.CommandExists("kaze") {
			systemupdate.Kaze(root)
		}
		if utils.CommandExists("snap") {
			systemupdate.Snap(root)
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
			fmt.Println("No package manager found!")
		}
	default:
		fmt.Println("Your system is currently not supported!")
	}
}
