package template

import (
	"archive/zip"
	"compress/gzip"
	"compress/zlib"
	"encoding/json"
	"os"

	"github.com/Stridsvagn69420/ganyu/utils"
)

func commandFile(pathtofile string, outputpath string) error {
	var cmd Command
	command, err := utils.ReadFileByteArray(pathtofile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(command, &cmd)
	if err != nil {
		return err
	}
	return utils.RunShell(false, cmd.Cmd, cmd.Args...)
}

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

func tarxzArchive(pathtofile string, outputpath string) error {
	file, err := os.Open(pathtofile)
	if err != nil {
		return err
	}
	defer file.Close()

	zl, err := zlib.NewReader(file)
	if err != nil {
		return err
	}
	defer zl.Close()

	return untar(zl, outputpath)
}

func targzArchive(pathtofile string, outputpath string) error {
	file, err := os.Open(pathtofile)
	if err != nil {
		return err
	}
	defer file.Close()

	gz, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gz.Close()

	return untar(gz, outputpath)
}
