package main

import (
	"os"

	"github.com/Stridsvagn69420/ganyu/utils"
	"github.com/Stridsvagn69420/pringo"
)

func main() {
	// Initialize command-line flags and read config
	config, err := config()
	if err != nil {
		utils.Printer.Errorln("Config file isn't present!", pringo.RedBright)
		utils.Printer.Errorln("Please read the instructions at "+REPOSITORY, pringo.Yellow)
		os.Exit(1)
	}

	// Get command
	switch os.Args[1] {
	case "update":
		sysupdate(OSType(), config.Sysupdate.Root, config.Sysupdate.CrossPkg)

	default:
		utils.Printer.Errorln("Command not found or doesn't exist yet!", pringo.Red)
		os.Exit(1)
	}

	os.Exit(0)
}
