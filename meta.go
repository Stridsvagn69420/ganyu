package main

import (
	"runtime"

	"github.com/Stridsvagn69420/pringo"
)

const (
	NAME       = "Ganyu"
	VERSION    = "v0.1.2"
	REPOSITORY = "https://github.com/Stridsvagn69420/ganyu"
	AUTHOR     = "Stridsvagn69420 (https://github.com/Stridsvagn69420)"
	LICENSE    = "GPL-3.0"
)

func PrintInfo() {
	// Meta
	cli.Println("Ganyu - Enhance your workflow across Linux distros and Windows", pringo.CyanBright)
	metaEntry("Version", VERSION)
	metaEntry("Author", AUTHOR)
	metaEntry("Repository", REPOSITORY)
	metaEntry("License", LICENSE)
	// System info
	metaEntry("OS", string(OSType()))
	metaEntry("Arch", runtime.GOARCH)
	metaEntry("Go", runtime.Version())
}

func metaEntry(key string, value string) {
	cli.Print(key+": ", pringo.Cyan)
	cli.Writeln(value)
}
