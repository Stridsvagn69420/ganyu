package main

import (
	"os"

	"github.com/Stridsvagn69420/ganyu/utils"
	"github.com/Stridsvagn69420/pringo"

	"github.com/Stridsvagn69420/ganyu/custom"
)

func main() {
	// Initialize command-line flags and read config
	config, err := config()
	if err != nil {
		utils.Printer.Errorln("Config file isn't present!", pringo.RedBright)
		utils.Printer.Errorln("Please read the instructions at "+REPOSITORY, pringo.YellowBright)
		os.Exit(1)
	}
	customcmd, err := custom.ReadCustom(custom.CustomCommandPath)
	if err != nil {
		utils.Printer.Errorln("Custom command file isn't present!", pringo.Yellow)
	}
	if len(os.Args) < 1 {
		// Error message
		// Print help
		// Exit 1
	} else {
		// Get command
		switch os.Args[1] {
		case "update":
			sysupdate(OSType(), config.Sysupdate.Root, config.Sysupdate.CrossPkg)

		default:
			// Try running custom command
			cmd, found := custom.FindCustom(os.Args[1], customcmd)
			if found {
				err := custom.RunCustom(cmd)
				if err != nil {
					utils.Printer.Errorln(err.Error(), pringo.Red)
				}
			} else {
				utils.Printer.Errorln("Command "+os.Args[1]+" not found or doesn't exist yet!", pringo.Red)
				os.Exit(1)
			}
		}
	}

	os.Exit(0)
}
