package template

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Stridsvagn69420/ganyu/utils"
	"github.com/Stridsvagn69420/pringo"
)

type Command struct {
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
}

type Archive string

const (
	CommandFile Archive = ".command"
	ZipArchive  Archive = ".zip"
	PwshScript  Archive = ".ps1"
	ShScript    Archive = ".sh"
	Invalid     Archive = ""
)

func GetArchiveType(pathfile string) Archive {
	ext := filepath.Ext(pathfile)
	switch ext {
	case ".zip":
		return ZipArchive
	case ".ps1":
		return PwshScript
	case ".sh":
		return ShScript
	case ".command":
		return CommandFile
	default:
		return Invalid
	}
}

func RunTemplate(tmplt Template, outputdir string) error {
	switch tmplt.Type {
	case ZipArchive:
		return zipArchive(tmplt.Path, outputdir)
	case PwshScript:
		return pwshFile(tmplt.Path, outputdir)
	case ShScript:
		return shFile(tmplt.Path, outputdir)
	case CommandFile:
		return commandFile(tmplt.Path, outputdir)
	}
	return nil
}

func findTemplate(list []Template, name string) (Template, error) {
	for _, tmplt := range list {
		if tmplt.Name == name {
			return tmplt, nil
		}
	}
	return Template{}, errors.New("Template not found")
}

func CreateTemplate(tmpl Template, dir string) error {
	if abspath, err := filepath.Abs(dir); err == nil {
		return RunTemplate(tmpl, abspath)
	} else {
		return RunTemplate(tmpl, dir)
	}
}

func GetTemplate(templs []Template, name string) Template {
	tmpl, err := findTemplate(templs, name)
	if err != nil {
		utils.Printer.Errorln("Template ", pringo.Red)
		utils.Printer.Error(name, pringo.RedBright)
		utils.Printer.Errorln(" not found!", pringo.Red)
		os.Exit(1)
	}
	return tmpl
}

func FinishTemplate(err error) {
	if err != nil {
		utils.Printer.Errorln("Template not created successfully!", pringo.Red)
		os.Exit(1)
	} else {
		utils.Printer.Println("Template created successfully!", pringo.Green)
	}
}
