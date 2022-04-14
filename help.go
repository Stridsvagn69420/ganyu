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

func PrintHelp(cstm []custom.Custom) {
	// ---- Print usage ----
	cli.Print("USAGE: ", pringo.CyanBright)
	cli.Println("ganyu <command> [args...]", pringo.White)
	cli.Writeln("")

	// ---- Print available commands ----
	cli.Println("COMMANDS:", pringo.CyanBright)

	// Built-in commands
	printCommand("update", "Updates Ganyu and the system via the package manager(s) available on your system")
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
