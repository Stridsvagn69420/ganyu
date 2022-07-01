package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/Stridsvagn69420/ganyu/custom"
	"github.com/Stridsvagn69420/ganyu/template"
	"github.com/Stridsvagn69420/ganyu/utils"
	"github.com/Stridsvagn69420/ganyu/ytdl"

	"github.com/Stridsvagn69420/pringo"
)

func main() {
	// Read config and init RPC
	config, err := config()
	if err != nil {
		cli.Errorln("Config file isn't present!", pringo.RedBright)
		cli.Errorln("Please read the instructions at "+REPOSITORY, pringo.YellowBright)
		os.Exit(1)
	}

	// Custom commands
	customcmd, err := custom.ReadCustom(custom.CustomCommandPath)
	if err != nil {
		cli.Errorln("Custom command file isn't present!", pringo.Yellow)
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
			switch sysupdate(OSType(), config.Sysupdate.Root, config.Sysupdate.CrossPkg) {
			case nil:
				cli.Println("System updated!", pringo.GreenBright)
			case ErrNotSupported:
				cli.Errorln("Your system is currently not supported!", pringo.Red)
				os.Exit(1)
			default:
				cli.Errorln("Failed to update system!", pringo.Red)
				os.Exit(0)
			}

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

		case "template":
			archivedir := filepath.Join(utils.GetHomeDir(), ".config", "ganyu", "templates")
			if !utils.DirExists(archivedir) {
				cli.Errorln("Template directory isn't present!", pringo.Red)
				cli.Errorln("See "+REPOSITORY+"/wiki/Templates for more!", pringo.Yellow)
				os.Exit(1)
			}
			templates := template.ListTemplates(archivedir)

			switch len(os.Args) {
			case 2:
				// List templates
				if len(templates) == 0 {
					cli.Println("No template available yet!", pringo.BlueBright)
				} else {
					cli.Println("Available templates:", pringo.BlueBright)
					for i, template := range templates {
						cli.Print("["+fmt.Sprint(i+1)+"] => ", pringo.Cyan)
						cli.Writeln(template.Name)
					}
				}

			case 3:
				if config.RPC {
					StartRPC()
				}
				go UpdateRPC("With template "+os.Args[2], "Creating a new project")
				// Create template in current directory
				cwd, _ := os.Getwd()
				err := template.CreateTemplate(
					template.GetTemplate(templates, os.Args[2]),
					cwd,
				)
				if err != nil {
					cli.Errorln("Template not created successfully!", pringo.Red)
					os.Exit(1)
				}

			case 4:
				if config.RPC {
					StartRPC()
				}
				go UpdateRPC("With template "+os.Args[2], "Creating a new project")
				// Create template in specified directory
				err := template.CreateTemplate(
					template.GetTemplate(templates, os.Args[2]),
					os.Args[3],
				)
				if err != nil {
					cli.Errorln("Template not created successfully!", pringo.Red)
					os.Exit(1)
				}

			default:
				cli.Errorln("Too many arguments!", pringo.Red)
				os.Exit(1)
			}

		case "gitpull":
			var rpctime time.Time
			if config.RPC {
				StartRPC()
				rpctime = time.Now()
			}
			pullerrors := 0
			for _, repopath := range config.Git.Repos {
				if config.RPC {
					go UpdateRPCTime(repopath, "Pulling from origin", rpctime)
				}
				cli.Println("["+repopath+"] ", pringo.Cyan)
				cmderr := utils.RunShellCwd(filepath.Join(config.Git.BaseDir, repopath), false, "git", "pull")
				if cmderr != nil {
					cli.Errorln("Error while pulling from "+repopath, pringo.Red)
					pullerrors++
				}
			}
			switch pullerrors {
			case 0:
				cli.Println("Done!", pringo.GreenBright)
			default:
				cli.Printf("%d error(s) occured while pulling from %d repositories!\n", pringo.Red, pullerrors, len(config.Git.Repos))
				os.Exit(1)
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
				cli.Errorln("Command "+os.Args[1]+" not found or doesn't exist yet!", pringo.RedBright)
				os.Exit(1)
			}
		}
	}
	go StopRPC()
	os.Exit(0)
}
