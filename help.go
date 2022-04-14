package main

import (
	"github.com/Stridsvagn69420/ganyu/custom"
	"github.com/Stridsvagn69420/ganyu/utils"
	"github.com/Stridsvagn69420/pringo"
)

var cmd *pringo.Printer = utils.Printer

func printCommand(name string, desc string) {
	cmd.Println("  "+name, pringo.WhiteBright)
	cmd.Println("  => "+desc, pringo.White)
}

func PrintHelp(cstm []custom.Custom) {
	// ---- Print usage ----
	cmd.Print("USAGE: ", pringo.Cyan)
	cmd.Println("ganyu <command> [args...]", pringo.White)
	cmd.Writeln("")

	// ---- Print available commands ----
	cmd.Println("COMMANDS:", pringo.Cyan)

	// Built-in commands
	printCommand("update", "Updates Ganyu and the system via the package manager(s) available on your system")
	printCommand("help", "Prints this help message")
	printCommand("version", "Prints the version of Ganyu as well as other information")

	// Custom commands
	if len(cstm) > 0 {
		cmd.Writeln("")
	}
	for _, c := range cstm {
		printCommand(c.Name, c.Desc)
	}
}

func PrintInfo() {
	// Meta
	cmd.Println("Ganyu - Enhance your workflow across Linux distros and Windows", pringo.CyanBright)
	cmd.Println("License: "+LICENSE, pringo.Cyan)
	cmd.Println("Version: "+VERSION, pringo.Cyan)
	cmd.Println("Repository: "+REPOSITORY, pringo.Cyan)
	cmd.Println("Author: "+AUTHOR, pringo.Cyan)

}
