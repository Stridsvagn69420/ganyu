package main

import (
	"net/url"
	"os"

	"github.com/Stridsvagn69420/ganyu/custom"
	"github.com/Stridsvagn69420/ganyu/utils"
	"github.com/Stridsvagn69420/ganyu/ytdl"

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
			if config.RPC {
				StartRPC()
			}
			go UpdateRPC(OSRich(), "Updating the system")
			sysupdate(OSType(), config.Sysupdate.Root, config.Sysupdate.CrossPkg)

		case "ytdl":
			if len(os.Args) < 5 {
				cli.Errorln("Please provide a Media Type, URL and Output directory!", pringo.RedBright)
				cli.Errorln("See "+os.Args[0]+" help for more.", pringo.Red)
				os.Exit(1)
			} else {
				if config.RPC {
					StartRPC()
				}
				// Check if Hostname exists in config
				url, err := url.Parse(os.Args[3])
				if err != nil {
					cli.Errorln("Invalid URL!", pringo.RedBright)
					os.Exit(1)
				}
				// Set Discord status
				go UpdateRPC(os.Args[2], "Downloading a video from "+url.Host)
				location := -1
				for n, i := range config.Ytdl {
					if i.Website == url.Host {
						location = n
					}
				}
				// Download the video
				ytdl.DownloadHandle(location, config.Ytdl)
			}

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
				if config.RPC {
					StartRPC()
				}
				UpdateRPC(cmd.Name, "Running a custom command")
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
	go StopRPC()
	os.Exit(0)
}
