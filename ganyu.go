package main

import (
	"os"

	"github.com/Stridsvagn69420/ganyu/custom"
	"github.com/Stridsvagn69420/ganyu/utils"
	"github.com/Stridsvagn69420/pringo"
)

func main() {
	// Read config and init RPC
	config, err := config()
	if err != nil {
		utils.Printer.Errorln("Config file isn't present!", pringo.RedBright)
		utils.Printer.Errorln("Please read the instructions at "+REPOSITORY, pringo.YellowBright)
		os.Exit(1)
	}
	if config.RPC {
		go StartRPC()
	}

	// Custom commands
	customcmd, err := custom.ReadCustom(custom.CustomCommandPath)
	if err != nil {
		utils.Printer.Errorln("Custom command file isn't present!", pringo.Yellow)
	}
	if len(os.Args) < 2 {
		PrintInfo(true)
		cli.Writeln("")
		PrintHelp(customcmd, true)
		os.Exit(1)
	} else {
		// Get command
		switch os.Args[1] {
		case "update":
			go UpdateRPC(OSRich(), "Updating the system")
			sysupdate(OSType(), config.Sysupdate.Root, config.Sysupdate.CrossPkg)

		case "info":
			PrintInfo(false)
			cli.Writeln("")
			PrintHelp(customcmd, false)

		case "help", "-h", "--help":
			PrintHelp(customcmd, false)

		case "version", "-V", "--version":
			PrintInfo(false)

		default:
			// Try running custom command
			cmd, found := custom.FindCustom(os.Args[1], customcmd)
			if found {
				go UpdateRPC(cmd.Name, "Running a custom command")
				err := custom.RunCustom(cmd)
				if err != nil {
					os.Exit(1)
				}
			} else {
				utils.Printer.Errorln("Command "+os.Args[1]+" not found or doesn't exist yet!", pringo.RedBright)
				os.Exit(1)
			}
		}
	}
	StopRPC()
	os.Exit(0)
}
