package main

import (
	"github.com/Stridsvagn69420/ganyu/custom"
	"github.com/Stridsvagn69420/ganyu/utils"
	"github.com/Stridsvagn69420/pringo"
)

var cli *pringo.Printer = utils.Printer

// ----- Help -----
func printCommand(name string, desc string) {
	cli.Println("  "+name, pringo.WhiteBright)
	cli.Println("  => "+desc, pringo.White)
}

func PrintHelp(cstm []custom.Custom, err bool) {
	var title pringo.Color
	if err {
		title = pringo.RedBright
	} else {
		title = pringo.CyanBright
	}
	// ---- Print usage ----
	cli.Print("USAGE: ", title)
	cli.Println("ganyu <command> [args...]", pringo.White)
	cli.Writeln("")

	// ---- Print available commands ----
	cli.Println("COMMANDS:", title)

	// Built-in commands
	printCommand("update", "Updates Ganyu and the system via the package manager(s) available on your system")
	printCommand("ytdl <URL> [<audio/video/combined> <Output>]", "Downloads a video using youtube-dl or yt-dlp")
	printCommand("help", "Prints this help message")
	printCommand("version", "Prints the version of Ganyu as well as other information")

	// Custom commands
	if len(cstm) > 0 {
		cli.Writeln("")
	}
	for _, c := range cstm {
		printCommand(c.Name, c.Desc)
	}
}
