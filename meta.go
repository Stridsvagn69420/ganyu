package main

import (
	"runtime"

	"github.com/Stridsvagn69420/pringo"
)

const (
	NAME       = "Ganyu"
	VERSION    = "v0.1.8"
	REPOSITORY = "https://github.com/Stridsvagn69420/ganyu"
	AUTHOR     = "Stridsvagn69420 (https://github.com/Stridsvagn69420)"
	LICENSE    = "GPL-3.0"
)

func PrintInfo(err bool) {
	var title pringo.Color
	var font pringo.Color
	if err {
		title = pringo.RedBright
		font = pringo.Red
	} else {
		title = pringo.CyanBright
		font = pringo.Cyan
	}
	// Meta
	cli.Println("Ganyu - Enhance your workflow across Linux distros and Windows", title)
	metaEntry("Version", VERSION, font)
	metaEntry("Author", AUTHOR, font)
	metaEntry("Repository", REPOSITORY, font)
	metaEntry("License", LICENSE, font)
	// System info
	metaEntry("OS", string(OSType()), font)
	metaEntry("Arch", runtime.GOARCH, font)
	metaEntry("Go", runtime.Version(), font)
}

func metaEntry(key string, value string, color pringo.Color) {
	cli.Print(key+": ", color)
	cli.Writeln(value)
}
