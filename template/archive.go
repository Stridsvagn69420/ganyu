package template

import (
	"archive/zip"

	"github.com/Stridsvagn69420/ganyu/utils"
)

func pwshFile(pathtofile string, outputpath string) error {
	if utils.CommandExists("pwsh") {
		return utils.RunShell(false, "pwsh", "-File", pathtofile)
	}
	return utils.RunShell(false, "powershell", "-File", pathtofile)
}

func shFile(pathtofile string, outputpath string) error {
	return utils.RunShell(false, "sh", pathtofile)
}

func zipArchive(pathtofile string, outputpath string) error {
	reader, err := zip.OpenReader(pathtofile)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, f := range reader.File {
		err := unzipFile(f, outputpath)
		if err != nil {
			return err
		}
	}
	return nil
}
