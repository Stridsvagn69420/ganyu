package template

import (
	"path/filepath"
	"strings"

	"github.com/Stridsvagn69420/ganyu/utils"
)

type Template struct {
	Name string
	Type Archive
	Path string
}

func ListTemplates(dir string) []Template {
	files := utils.ReadFiles(dir)
	var templates []Template
	for _, file := range files {
		switch filepath.Ext(file) {
		case ".zip":
			templates = append(templates, newTemplate(file, ".zip", ZipArchive, dir))
		case ".ps1":
			templates = append(templates, newTemplate(file, ".ps1", PwshScript, dir))
		case ".sh":
			templates = append(templates, newTemplate(file, ".sh", ShScript, dir))
		case ".command":
			templates = append(templates, newTemplate(file, ".command", CommandFile, dir))
		}
	}
	return templates
}

func newTemplate(file string, ext string, archtype Archive, dir string) Template {
	return Template{
		Name: strings.ReplaceAll(file, ext, ""),
		Type: ZipArchive,
		Path: filepath.Join(dir, file),
	}
}
