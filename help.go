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
	// Base commands in the order they got added
	printCommand("update", "Updates Ganyu and the system via the package manager(s) available on your system")
	printCommand("ytdl <audio/video/combined> <URL> <Output>", "Downloads a video using youtube-dl or yt-dlp")
	printCommand("template <TemplateName> [Outdir]", "Creates a project from a template file, e.g. a script or an archive, in the current directory unless specified.")
	printCommand("gitpull", "Pull from all repositories listed in the config file")
	printCommand("build", "Runs the Ganyu Build System to compile the current project")
	// Help utils
	printCommand("help", "Prints this help message")
	printCommand("version", "Prints the version of Ganyu as well as other information")
	printCommand("info", "Prints both version and help message")

	// Custom commands
	if len(cstm) > 0 {
		cli.Writeln("")
	}
	for _, c := range cstm {
		printCommand(c.Name, c.Desc)
	}
}
